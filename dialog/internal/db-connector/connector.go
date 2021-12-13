package dbconn

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/m-shev/otus-social/dialog/internal/configuration"
	"log"
	"time"
)

type Connector struct {
	dbs map[configuration.DbId]*sql.DB
}

func NewDbConnector() *Connector {
	return &Connector{dbs: make(map[configuration.DbId]*sql.DB)}
}

func (c *Connector) GetConnection(id configuration.DbId) (*sql.DB, error) {
	db, ok := c.dbs[id]

	if !ok {
		textErr := fmt.Sprintf(
			"Db connection with name: %s was not create. Use AddDbList to create connection.", id)
		return nil, errors.New(textErr)
	}

	return db, nil
}

func (c *Connector) AddDb(conf configuration.DbConfig) error {
	return c.AddDbList([]configuration.DbConfig{conf})
}

func (c *Connector) AddDbList(list []configuration.DbConfig) error {
	for _, v := range list {
		db, err := c.createConnection(v)

		if err != nil {

			log.Println(errorCannotCreateConn(v.DbId, err))

			return err
		}

		c.dbs[v.DbId] = db
	}

	return nil
}

func (c *Connector) createConnection(conf configuration.DbConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", c.dbUrl(conf))

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * conf.ConnMaxLifetime)
	db.SetMaxOpenConns(conf.MaxOpenConnection)
	db.SetMaxIdleConns(conf.MaxIdleConnection)

	return db, nil
}

func (c *Connector) dbUrl(conf configuration.DbConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		conf.User, conf.Password, conf.Host, conf.Port, conf.DbName)
}

func errorCannotCreateConn(id configuration.DbId, err error) string {
	return fmt.Sprintf("Received an error while trying to create a connection for: %s, error: %s",
		id, err.Error())
}
