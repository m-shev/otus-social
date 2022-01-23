package config

import "time"

type Config struct {
	Db
	Server
	Protection
	Broker
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

type Broker struct {
	BrokerUrls []string
	PostTopic  Topic
}

type Topic struct {
	Name              string
	NumPartitions     int
	ReplicationFactor int
}
