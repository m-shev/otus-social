package config

import "time"

type Config struct {
	Db
	Server
	Protection
}

type Db struct {
	Password          string
	User              string
	Host              string
	Port              string
	DbName            string
	MaxOpenConnection int
	MaxIdleConnection int
	ConnMaxLifetime   time.Duration
	MigrationPath     string
}

type Server struct {
	StartDelay time.Duration
	Host       string
	Port       string
}

type Protection struct {
	AllowOrigins []string
}
