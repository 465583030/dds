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

var users *friends.Users
var frands *friends.Friends

func main() {

	const (
		defaultLinuxConfigPath   = "/usr/local/services/dds/conf/config.json"
		defaultWindowsConfigPath = "C:\\Program Files\\dds\\conf\\config.json"
		usage                    = "\"/usr/local/services/dds/conf/config.json\""
	)

	defaultConfigPath := ""
	if runtime.GOOS == "windows" { //windows
		defaultConfigPath = defaultWindowsConfigPath
	} else { //linux
		defaultConfigPath = defaultLinuxConfigPath
	}

	configPath := flag.String("conf", defaultConfigPath, usage)
	flag.Parse()

	cfg := &config.Config{}
	if err := cfg.ReadConfig(*configPath); err != nil {
		log.Fatal(err)
	}

	// Users & Friends
	users = &friends.Users{}
	usr := friends.User{
		Username: "ricl",
	}
	(*users)[usr.Username] = usr

	frands = &friends.Friends{}

	// GRPC
	go routines.MainRoutine(users, frands, cfg)

	// file server
	http.Handle("/", http.FileServer(http.Dir(cfg.Directory)))

	// API
	apiHandler, err := handler.CreateAPIHandler(cfg)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/api/", apiHandler)

	log.Println("\ndds started using config file of", *configPath, "\nwith parameters", cfg.ToString())
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	go log.Fatal(http.ListenAndServe(addr, nil))

}
