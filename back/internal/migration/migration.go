package migration

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/m-shev/otus-social/back/internal/config"
	"log"
)

type Manager struct {
	dbConf config.Db
	db     *sql.DB
	Logger *log.Logger
}

func NewManager(conf config.Db, logger *log.Logger) *Manager {
	return &Manager{dbConf: conf, Logger: logger}
}

func (manager *Manager) Up() {
	defer recovery()
	if err := manager.create().Up(); err != nil {
		manager.handleMigrationErr(err)
	} else {
		manager.Logger.Println("Migration executed success")
	}

	manager.close()
}

func (manager *Manager) Down() {
	defer recovery()
	manager.handleMigrationErr(manager.create().Down())
	manager.close()
}

func (manager *Manager) Force(version int) {
	defer recovery()
	manager.handleMigrationErr(manager.create().Force(version))
	manager.close()
}

func (manager *Manager) create() *migrate.Migrate {
	driver, err := mysql.WithInstance(manager.open(), &mysql.Config{})
	manager.handleMigrationErr(err)

	m, err := migrate.NewWithDatabaseInstance(
		manager.migrationPath(),
		"mysql",
		driver,
	)

	manager.handleMigrationErr(err)

	return m
}

func (manager *Manager) open() *sql.DB {
	if manager.db == nil {
		db, err := sql.Open("mysql", manager.dbUrl())
		manager.handleMigrationErr(err)
		manager.db = db
	}

	return manager.db
}

func (manager *Manager) close() {
	if manager.db != nil {
		err := manager.db.Close()
		manager.handleMigrationErr(err)
	}
}

func (manager *Manager) migrationPath() string {
	return fmt.Sprintf("file://%s", manager.dbConf.MigrationPath)
}

func (manager *Manager) dbUrl() string {
	conf := manager.dbConf

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		conf.User, conf.Password, conf.Host, conf.Port, conf.DbName)
}

func recovery() {
	_ = recover()
}

func (manager *Manager) handleMigrationErr(err error) {
	if err != nil {
		if err.Error() == "no change" {
			manager.Logger.Println("Migration: no change")
		} else {
			manager.Logger.Println("Migration error: ", err)
			panic("")
		}
	}
}
