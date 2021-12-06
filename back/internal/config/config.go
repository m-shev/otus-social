package config

import (
	"github.com/m-shev/go-config"
	"github.com/spf13/viper"
	"log"
)

var EnvType = struct {
	Dev  string
	QA   string
	Prod string
}{Dev: "DEV", Prod: "PROD", QA: "QA"}

var conf = &Config{}

var configFiles = map[string]string{
	EnvType.Dev:  "dev",
	EnvType.QA:   "qa",
	EnvType.Prod: "prod",
}

const (
	prefix            = "social"
	envVar            = "env"
	defaultConfig     = "default"
	defaultConfigPath = "./config"
)

var goConfig = goconfig.NewGoConfig(goconfig.Option{
	Prefix:            prefix,
	Config:            conf,
	DefaultConfig:     defaultConfig,
	DefaultConfigPath: defaultConfigPath,
	ConfigFiles:       configFiles,
})

func GetConfig() Config {
	i := goConfig.GetConfig()
	c, ok := i.(Config)

	if !ok {
		log.Fatal("Cannot cast config to type Config")
	}

	return c
}

//
//func GetEnv() string {
//	return conf.env
//}
//
//func AddConfigPath(path string) {
//	viper.AddConfigPath(path)
//}
//
//func readConfig() {
//	defineEnv()
//	readDefault()
//	readTargetConfig()
//	conf.isRead = true
//}
//
//func readDefault() {
//	viper.AddConfigPath(defaultConfigPath)
//	viper.SetConfigName(defaultConfig)
//
//	read()
//	unmarshal()
//}
//
//func readTargetConfig() {
//	configName, ok := configFiles[conf.env]
//
//	if ok {
//		viper.SetConfigName(configName)
//		read()
//		unmarshal()
//	} else {
//		log.Fatal("Cannot read target config", configName)
//	}
//}
//
//func defineEnv() {
//	viper.AutomaticEnv()
//	viper.SetEnvPrefix(prefix)
//
//	env := viper.GetString(envVar)
//
//	switch env {
//	case EnvType.Prod:
//		conf.env = EnvType.Prod
//	case EnvType.QA:
//		conf.env = EnvType.QA
//	default:
//		conf.env = EnvType.Dev
//	}
//}
//
func unmarshal() {
	err := viper.Unmarshal(conf)

	if err != nil {
		log.Fatal(err)
	}
}

func read() {
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}
}
