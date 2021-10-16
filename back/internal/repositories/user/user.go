package user

import (
	"context"
	"database/sql"
	"github.com/m-shev/otus-social/internal/connector"
	"time"
)

const queryTimeout = time.Second * 5

type Repository struct {
	db *sql.DB
}

func NewRepository(con *connector.Connector) *Repository {
	return &Repository{db: con.GetConnection()}
}

func (r *Repository) Create(form CreateUserForm) (User, error) {
	query := `insert into user(name, surname, age, city, email, password)
				values (?, ?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(r.createContext(),
		query, form.Name, form.Surname, form.Age, form.City, form.Email, form.Password)

	if err != nil {
		return User{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return User{}, err
	}

	return r.GetById(int(id))
}

func (r *Repository) GetById(id int) (User, error) {
	user := User{}

	query := `select * from user where id=?`

	row := r.db.QueryRowContext(r.createContext(), query, id)

	err := row.Scan(&user.id, &user.Name, &user.Surname, &user.Age, &user.City, &user.Email, &user.Password)

	return user, err
}

func (r *Repository) createContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), queryTimeout)
	return ctx
}