package util

import (
	"log"
	"net/url"
	"os"

	"github.com/BurntSushi/toml"
)

const ConfigPath string = "../config.toml"

var DefaultConfig Config

type Config struct {
	Network  NetworkConfig  `toml:"network"`
	Database DatabaseConfig `toml:"database"`
	Bot      BotConfig      `toml:"bot"`
}

type NetworkConfig struct {
	Proxy string
}

type DatabaseConfig struct {
	Filepath string
}

type BotConfig struct {
	Token string
}

func init() {
	// 读取并初始化config
	if _, err := os.Stat(ConfigPath); err != nil {
		log.Fatalln(0, err)
	}
	if _, err := toml.DecodeFile(ConfigPath, &DefaultConfig); err != nil {
		log.Fatalln(0, err)
	}
	if (DefaultConfig.Bot == BotConfig{} || DefaultConfig.Bot.Token == "") {
		log.Fatalln("no bot token detected")
	}
}

func (config *Config) PrintConfig() {
	if !config.Network.NeedProxy() {
		log.Println("no proxy config detected, using direct")
	} else {
		log.Printf("proxy %s\n", config.Network.Proxy)
	}
}

func (networkConfig *NetworkConfig) NeedProxy() bool {
	return *networkConfig != NetworkConfig{} && networkConfig.Proxy != ""
}

func (networkConfig *NetworkConfig) GetProxyProtocal() string {
	u, err := url.Parse(networkConfig.Proxy)
	if err != nil {
		log.Fatalln(0, err)
	}
	return u.Scheme
}
