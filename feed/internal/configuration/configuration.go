package configuration

import goconfig "github.com/m-shev/go-config"

const (
	prefix            = "feed"
	defaultConfig     = "default"
	defaultConfigPath = "./configuration"
)

var conf = &Configuration{}

var configFiles = map[string]string{
	goconfig.DEV:  "dev",
	goconfig.QA:   "qa",
	goconfig.PROD: "prod",
}

var isRead bool = false

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

		isRead = true
	}

	return *conf
}
