package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

// Config download master's common config
type Config struct {
	Host      string `json:"host"`
	Username  string `json:"username"`
	Port      int    `json:"port"`
	RPCPort   int    `json:"rpc_port"`
	Token     string `json:"token"`
	Directory string `json:"directory"`
	Location  string `json:"location"`
}

// ReadConfig read config from specified file
func (cfg *Config) ReadConfig() error {
	readBytes, err := ioutil.ReadFile(cfg.Location)
	if err != nil {
		return err
	}
	err = json.Unmarshal(readBytes, cfg)
	if err != nil {
		return err
	}
	return nil
}

// WriteConfig write config in memory to specified file
func (cfg *Config) WriteConfig() error {
	payload, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(cfg.Location, payload, 0644)
	if err != nil {
		return err
	}
	return nil
}

// ToString format cfg to a string
func (cfg *Config) ToString() string {
	apiAddrString := "127.0.0.1:" + strconv.Itoa(cfg.Port)
	return fmt.Sprintf("api-addr: [%s], rpc-addr: [%s:%d] username: [%s], token: [%s], working directory: [%s], confg location:[%s]", apiAddrString, cfg.Host, cfg.RPCPort, cfg.Username, cfg.Token, cfg.Directory, cfg.Location)
}
