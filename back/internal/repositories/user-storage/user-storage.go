package user_storage

import (
	"context"
	"database/sql"
	"github.com/m-shev/otus-social/internal/connector"
	"github.com/m-shev/otus-social/internal/services/user"
	"strings"
	"time"
)

const queryTimeout = time.Second * 5

type Repository struct {
	db *sql.DB
}

func NewRepository(con *connector.Connector) *Repository {
	return &Repository{db: con.GetConnection()}
}

func (r *Repository) Create(form user.CreateUserForm) (user.User, error) {
	var u user.User
	query := `insert into user(name, surname, age, city, email, password)
				values (?, ?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(r.createContext(),
		query, form.Name, form.Surname, form.Age, form.City, form.Email, form.Password)

	if err != nil {
		if strings.Contains(err.Error(), sqlDuplicateEntry) {
			return u, user.ErrorUserAlreadyCreated
		}
		return u, err
	}

	id, err := result.LastInsertId()
	userId := int(id)

	if err != nil {
		return u, err
	}

	err = r.addInterest(userId, form.Interests)

	if err != nil {
		return u, err
	}

	return r.GetById(userId)
}

func (r *Repository) GetById(id int) (user.User, error) {
	u := user.User{}

	query := `select * from user where id=?`

	row := r.db.QueryRowContext(r.createContext(), query, id)

	err := row.Scan(&u.Id, &u.Name, &u.Surname, &u.Age, &u.City, &u.Email, &u.Password)

	if err != nil {
		return u, err
	}

	interests, err := r.getUserInterests(u.Id)

	u.Interests = interests

	return u, err
}

func (r *Repository) FindUser(email string) (user.User, error) {
	u := user.User{}
	query := `select * from user where email=?`

	row := r.db.QueryRowContext(r.createContext(), query, email)
	err := row.Scan(&u.Id, &u.Name, &u.Surname, &u.Age, &u.City, &u.Email, &u.Password)

	if err != nil {
		if err.Error() == sqlNotFoundError {
			return user.User{}, user.ErrorUserNotFound
		}

		return user.User{}, err
	}

	interests, err := r.getUserInterests(u.Id)

	u.Interests = interests

	return u, nil
}

func (r * Repository) getUserInterests(userId int) ([]user.Interest, error) {
	query := `select interest.id, interest.name from interest left join user_interest on interestId=id
				where userId=?`

	rows, err := r.db.QueryContext(r.createContext(), query, userId)

	if err != nil {
		return nil, err
	}

	interests := make([]user.Interest, 0)

	for rows.Next() {
		var interest user.Interest

		if err = rows.Scan(&interest.Id, &interest.Name); err != nil {
			return nil, err
		}

		interests = append(interests, interest)
	}

	return interests, err
}

func (r *Repository) addInterest(userId int, interestNames []string) error {
	if len(interestNames) == 0 {
		return nil
	}

	interests, err := r.getInterestsByNames(interestNames)

	if err != nil {
		return err
	}

	query := "insert into user_interest(userId, interestId) values"

	params := make([]interface{}, 0)

	for i, v := range interests {
		if i ==0 {
			query += " (?, ?)"
		} else {
			query += ", (?, ?)"
		}

		params = append(params, userId)
		params = append(params, v.Id)
	}

	_, err = r.db.ExecContext(r.createContext(), query, params...)

	return err
}



func (r *Repository) createContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), queryTimeout)
	return ctx
}
