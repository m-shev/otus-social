package user_storage

import (
	"github.com/m-shev/otus-social/back/internal/services/user"
	"strings"
)

func (r *Repository) getInterestsByNames(interests []string) ([]user.Interest, error) {
	result := make([]user.Interest, 0)

	if len(interests) == 0 {
		return result, nil
	}

	found, err := r.findInterestsByNames(interests)

	if err != nil {
		return result, err
	}

	newInterests := notExistedInterests(found, interests)

	created, err := r.createInterests(newInterests)

	if err != nil {
		return result, err
	}

	for _, v := range found {
		result = append(result, v)
	}

	result = append(result, created...)

	return result, nil
}

func notExistedInterests(existed map[string]user.Interest, interests []string) []string {
	newInterest := make([]string, 0)

	for _, v := range interests {
		if _, ok := existed[v]; !ok {
			newInterest = append(newInterest, v)
		}
	}

	return newInterest
}

func (r *Repository) createInterests(interest []string) ([]user.Interest, error) {
	created := make([]user.Interest, 0)

	if len(interest) == 0 {
		return created, nil
	}

	query := "insert into interest(name) values "
	params := make([]interface{}, 0)

	for i, v := range interest {
		if i == 0 {
			query += " (?)"
		} else {
			query += ", (?)"
		}

		params = append(params, v)
	}

	_, err := r.db.ExecContext(r.createContext(), query, params...)

	if err != nil {
		return created, err
	}

	found, err := r.findInterestsByNames(interest)

	if err != nil {
		return created, err
	}

	for _, v := range found {
		created = append(created, v)
	}

	return created, nil
}

func (r *Repository) findInterestsByNames(interests []string) (map[string]user.Interest, error) {
	query := "select * from interest where name in (?" + strings.Repeat(",?", len(interests)-1) + ")"
	args := make([]interface{}, 0, len(interests))
	found := make(map[string]user.Interest)

	for _, v := range interests {
		args = append(args, v)
	}

	rows, err := r.db.QueryContext(r.createContext(), query, args...)

	for rows.Next() {
		interest := user.Interest{}

		if err = rows.Scan(&interest.Id, &interest.Name); err != nil {
			return found, err
		}

		found[interest.Name] = interest
	}

	return found, err
}
