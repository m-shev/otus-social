package configuration

import "time"

type Configuration struct {
	Server
	Db
	Protection
	Broker
	Topic
}

type Server struct {
	Host string
	Port string
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
