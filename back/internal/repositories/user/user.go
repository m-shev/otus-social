package user

import (
	"database/sql"
	"github.com/m-shev/otus-social/internal/connector"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(con *connector.Connector) *Repository {
	return &Repository{db: con.GetConnection()}
}

//func (r *Repository) Create(form CreateUserForm) User {
//	query := `insert into table user(name, surname, age, city, email, password)
//				values ($1, $2, $3, $4, $5, $6)`
//	r.db.Exec()
//}