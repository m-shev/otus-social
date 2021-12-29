package config

import (
	"github.com/m-shev/go-config"
)

var conf = &Config{}

var configFiles = map[string]string{
	goconfig.DEV:  "dev",
	goconfig.QA:   "qa",
	goconfig.PROD: "prod",
}

const (
	prefix            = "social"
	defaultConfig     = "default"
	defaultConfigPath = "./config"
)

var isRead bool

func GetConfig() Config {

	if !isRead {
		con := goconfig.NewGoConfig(goconfig.Option{
			Prefix:            prefix,
			Config:            conf,
			DefaultConfig:     defaultConfig,
			DefaultConfigPath: defaultConfigPath,
			ConfigFiles:       configFiles,
		})

		err := con.ReadConfig()

		if err != nil {
			panic(err.Error())
		}
	}

	return *conf
}
