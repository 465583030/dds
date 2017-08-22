package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/riclava/dds/api/handler"
	"github.com/riclava/dds/cluster/config"
	"github.com/riclava/dds/cluster/friends"
	"github.com/riclava/dds/cluster/routines"
)

var friands *friends.Friends

const (
	defaultLinuxConfigPath   = "/usr/local/services/dds/conf/config.json"
	defaultWindowsConfigPath = "C:\\Program Files\\dds\\conf\\config.json"
	usage                    = "\"/usr/local/services/dds/conf/config.json\""

	defaultLinuxFriendsConfigPath   = "/usr/local/services/dds/conf/friends.json"
	defaultWindowsFriendsConfigPath = "C:\\Program Files\\dds\\conf\\friends.json"
	usageFriends                    = "\"/usr/local/services/dds/conf/friends.json\""
)

func loadConfig() (*config.Config, *friends.Friends) {
	defaultConfigPath := ""
	if runtime.GOOS == "windows" { //windows
		defaultConfigPath = defaultWindowsConfigPath
	} else { //linux
		defaultConfigPath = defaultLinuxConfigPath
	}
	configPath := flag.String("conf", defaultConfigPath, usage)

	defaultConfigFriendsPath := ""
	if runtime.GOOS == "windows" { //windows
		defaultConfigFriendsPath = defaultWindowsFriendsConfigPath
	} else { //linux
		defaultConfigFriendsPath = defaultLinuxFriendsConfigPath
	}
	configFriendsPath := flag.String("friends", defaultConfigFriendsPath, usageFriends)

	flag.Parse()

	cfg := &config.Config{
		Location: *configPath,
	}
	if err := cfg.ReadConfig(); err != nil {
		log.Fatal(err)
	}

	friands = &friends.Friends{
		Location: *configFriendsPath,
	}
	if err := friands.ReadConfig(); err != nil {
		log.Fatal(err)
	}
	return cfg, friands
}

func main() {

	cfg, friands := loadConfig()

	myself := &friends.Friend{}
	myself.Username = cfg.Username
	myself.Host = cfg.Host
	myself.Port = cfg.RPCPort
	myself.Token = cfg.Token

	friands.Friends[myself.Username] = *myself

	// GRPC
	go routines.MainRoutine(friands, cfg)

	// Process task
	go routines.TaskProcessRoutine(cfg, myself, friands)

	// file server
	http.Handle("/", http.FileServer(http.Dir(cfg.Directory)))

	// API
	apiHandler, err := handler.CreateAPIHandler(cfg, friands)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/api/", apiHandler)

	log.Println("\ndds started using config file from", cfg.Location, "\nwith parameters", cfg.ToString())
	log.Println("\nloading friends config from", friands.Location)
	addr := fmt.Sprintf("%s:%d", "127.0.0.1", cfg.Port)
	go log.Fatal(http.ListenAndServe(addr, nil))

}
