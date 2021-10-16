package connector

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/m-shev/otus-social/internal/config"
	"log"
	"time"
)

type Connector struct {
	conf config.Db
	logger *log.Logger
	db *sql.DB
}


func NewDbConnector(conf config.Db, logger *log.Logger) *Connector {
	return &Connector{
		conf:   conf,
		logger: logger,
		db: nil,
	}
}

func (c *Connector) GetConnection() *sql.DB {

	if c.db != nil {
		return c.db
	}

	db, err := sql.Open("mysql", c.dbUrl())

	if err != nil {
		c.logger.Fatal("Cannot get db connection: ", err)
	}

	db.SetConnMaxLifetime(time.Minute * c.conf.ConnMaxLifetime)
	db.SetMaxOpenConns(c.conf.MaxOpenConnection)
	db.SetMaxIdleConns(c.conf.MaxIdleConnection)

	return db
}

func (c *Connector) dbUrl() string {
	return fmt.Sprintf("%s:%s@/%s", c.conf.User, c.conf.Password, c.conf.DbName)
}
