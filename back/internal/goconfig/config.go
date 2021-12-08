package goconfig

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
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
		prefix:            option.Prefix,
		env:               DEV,
		config:            option.Config,
		defaultConfig:     option.DefaultConfig,
		defaultConfigPath: option.DefaultConfigPath,
		configFiles:       option.ConfigFiles,
	}
}

func (c *GoConfig) ReadConfig() error {
	return c.readConfig()
}

func (c *GoConfig) GetEnv() string {
	return c.env
}

func (c *GoConfig) readConfig() error {
	c.defineEnv()

	err := c.readDefault()

	if err != nil {
		return err
	}

	err = c.readTargetConfig()

	return err
}

func (c *GoConfig) readDefault() error {
	viper.AddConfigPath(c.defaultConfigPath)
	viper.SetConfigName(c.defaultConfig)

	return c.unmarshal()
}

func (c *GoConfig) readTargetConfig() error {
	configName, ok := c.configFiles[c.env]

	if !ok {
		return errors.New(fmt.Sprintf("The target configuration with name %s cannot be read", configName))
	}

	viper.SetConfigName(configName)

	return c.unmarshal()
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

func (c *GoConfig) unmarshal() error {
	err := viper.ReadInConfig()

	if err != nil {
		return err
	}
	err = viper.Unmarshal(c.config)

	if err != nil {
		return err
	}

	return nil
}
