package config

// Config download master's common config
type Config struct {
	Addr  string `json:"addr"`
	Port  int    `json:"port"`
	Token string `json:"token"`
}
