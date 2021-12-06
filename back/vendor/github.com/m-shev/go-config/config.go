package goconfig

import (
	"github.com/spf13/viper"
	"log"
)

const (
	DEV  = "DEV"
	QA   = "QA"
	PROD = "PROD"
)

const (
	envVar = "env"
)

type GoConfig struct {
	isRead            bool
	prefix            string
	env               string
	config            interface{}
	defaultConfig     string
	defaultConfigPath string
	configFiles       map[string]string
}

type Option struct {
	Prefix            string
	Config            interface{}
	DefaultConfig     string
	DefaultConfigPath string
	ConfigFiles       map[string]string
}

func NewGoConfig(option Option) *GoConfig {
	return &GoConfig{
		isRead:            false,
		prefix:            option.Prefix,
		env:               DEV,
		config:            option.Config,
		defaultConfig:     option.DefaultConfig,
		defaultConfigPath: option.DefaultConfigPath,
	}
}

func (c *GoConfig) GetConfig() interface{} {
	if !c.isRead {
		c.readConfig()
	}

	return c.config
}

func (c *GoConfig) GetEnv() string {
	return c.env
}

func (c *GoConfig) readConfig() {
	c.defineEnv()
	c.readDefault()
	c.readTargetConfig()
	c.isRead = true
}

func (c *GoConfig) readDefault() {
	viper.AddConfigPath(c.defaultConfigPath)
	viper.SetConfigName(c.defaultConfig)

	c.read()
	c.unmarshal()
}

func (c *GoConfig) readTargetConfig() {
	configName, ok := c.configFiles[c.env]

	if ok {
		viper.SetConfigName(configName)
		c.read()
		c.unmarshal()
	} else {
		log.Fatal("Cannot read target config", configName)
	}
}

func (c *GoConfig) defineEnv() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(c.prefix)

	env := viper.GetString(envVar)

	switch env {
	case PROD:
		c.env = PROD
	case QA:
		env = QA
	default:
		c.env = DEV
	}
}

func (c *GoConfig) unmarshal() {
	err := viper.Unmarshal(c.config)

	if err != nil {
		log.Fatal(err)
	}
}

func (c *GoConfig) read() {
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}
}
