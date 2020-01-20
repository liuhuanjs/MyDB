package config

import "sync/atomic"

//server config
type Config struct {
	Host             string `toml:"host" json:"host"`
	AdvertiseAddress string `toml:"advertise-address" json:"advertise-address"`
	Port             uint   `toml:"port" json:"port"`
}

var defaultConf = Config{
	Host:             "127.0.0.1",
	AdvertiseAddress: "",
	Port:             6666,
}

var globalConf = atomic.Value{}

func GetGlobalConfig() *Config {
	//return globalConf.Load().(*Config)
	return &defaultConf
}
