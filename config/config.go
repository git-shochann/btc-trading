package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	ApiKey string
	ApiSecret string
	LogFile string
}

var Config ConfigList // グローバルで呼び出せるようにする

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("failed to read file %v", err)
		os.Exit(1) // プログラム終了
	}
	Config = ConfigList{
		cfg.Section("bitflyer").Key("api_key").String(),
		cfg.Section("bitflyer").Key("api_secret").String(),
		cfg.Section("gotrading").Key("log_file").String(),
    }
}