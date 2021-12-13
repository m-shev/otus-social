package migration

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"log"
)

type DbConfig struct {
	User          string
	Password      string
	Host          string
	Port          string
	DbName        string
	MigrationPath string
}

type Manager struct {
	db *sql.DB
}

func NewMigrationHelper() *Manager {
	return &Manager{}
}

func (m *Manager) Up(dbConf DbConfig) {
	defer recovery()
	if err := m.create(dbConf).Up(); err != nil {
		m.handleMigrationErr(err)
	} else {
		log.Println("Migration executed success")
	}

	m.close()
}

func (m *Manager) Down(dbConf DbConfig) {
	defer recovery()
	m.handleMigrationErr(m.create(dbConf).Down())
	m.close()
}

func (m *Manager) Force(version int, dbConfig DbConfig) {
	defer recovery()
	m.handleMigrationErr(m.create(dbConfig).Force(version))
	m.close()
}

func (m *Manager) create(dbConfig DbConfig) *migrate.Migrate {
	driver, err := mysql.WithInstance(m.open(dbConfig), &mysql.Config{})
	m.handleMigrationErr(err)

	migration, err := migrate.NewWithDatabaseInstance(
		m.migrationPath(dbConfig.MigrationPath),
		"mysql",
		driver,
	)

	m.handleMigrationErr(err)

	return migration
}

func (m *Manager) open(dbConfig DbConfig) *sql.DB {
	if m.db == nil {
		db, err := sql.Open("mysql", m.dbUrl(dbConfig))
		m.handleMigrationErr(err)
		m.db = db
	}

	return m.db
}

func (m *Manager) close() {
	if m.db != nil {
		err := m.db.Close()
		m.handleMigrationErr(err)
	}
}

func (m *Manager) migrationPath(path string) string {
	return fmt.Sprintf("file://%s", path)
}

func (m *Manager) dbUrl(dbConfig DbConfig) string {
	conf := dbConfig

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		conf.User, conf.Password, conf.Host, conf.Port, conf.DbName)
}

func recovery() {
	_ = recover()
}

func (m *Manager) handleMigrationErr(err error) {
	if err != nil {
		if err.Error() == "no change" {
			log.Println("Migration: no change")
		} else {
			log.Println("Migration error: ", err)
			panic("")
		}
	}
}
