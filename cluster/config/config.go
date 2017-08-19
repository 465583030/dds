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
}

// ReadConfig read config from specified file
func (cfg *Config) ReadConfig(path string) error {
	config, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(config, cfg)
	if err != nil {
		return err
	}
	return nil
}

// WriteConfig write config in memory to specified file
func (cfg *Config) WriteConfig(path string) error {
	payload, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, payload, 0644)
	if err != nil {
		return err
	}
	return nil
}

// ToString format cfg to a string
func (cfg *Config) ToString() string {
	addrString := cfg.Host + ":" + strconv.Itoa(cfg.Port)
	return fmt.Sprintf("add: [%s], rpc_port: [%d] username: [%s], token: [%s], working directory: [%s]", addrString, cfg.RPCPort, cfg.Username, cfg.Token, cfg.Directory)
}
