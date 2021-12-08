package configuration

import (
	"github.com/m-shev/go-config"
	"time"
)

type Configuration struct {
	Server
}

type Server struct {
	StartDelay time.Duration
	Host       string
	Port       string
}

const (
	prefix            = "dialog"
	defaultConfig     = "default"
	defaultConfigPath = "./config"
)

var conf = &Configuration{}

var configFiles = map[string]string{
	goconfig.DEV:  "dev",
	goconfig.QA:   "qa",
	goconfig.PROD: "prod",
}

var isRead bool

func GetConfig() Configuration {

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
