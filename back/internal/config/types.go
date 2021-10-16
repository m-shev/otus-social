package config

import "time"

type Config struct {
	Db
	isRead bool
	env    string
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

type App struct {
	Addr               string
	CacheFlushInterval time.Duration
	Protection
}

type Protection struct {
	AllowHosts []string
	DevToken   string
}

type Api struct {
	RestoreCache       bool
	CacheSize          int64
	CachePath          string
	DownloadTimeout    time.Duration
	FileFolder         string
	CacheControlHeader string
}

type Cookie struct {
	Key   string
	Value string
}

type Logging struct {
	ErrorLog  LogParams
	AccessLog LogParams
}

type LogParams struct {
	FileName   string
	MaxBackups int
	MaxAge     int
}
