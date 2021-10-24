package user_storage

import (
	"database/sql"
	"fmt"
	"github.com/m-shev/otus-social/internal/services/user"
)

const selectStr = "select user.id, user.name, user.surname, user.avatar from user"

func (r *Repository) FindUsers(form user.FindUsersForm) ([]user.Friend, error) {
	var err error
	var rows *sql.Rows

	if form.Name != "" && form.Surname != "" {
		rows, err = r.fullUserQuery(form)
	} else if form.Name != "" {
		rows, err = r.findUserByQuery(form, "name")
	} else if form.Surname != "" {
		rows, err = r.findUserByQuery(form, "surname")
	} else {
		rows, err = r.findUserQuery(form)
	}

	if err != nil {
		return nil, err
	}

	users := make([]user.Friend, 0)

	for rows.Next() {
		var friend user.Friend

		if err = scanFriend(rows, &friend); err != nil {
			return nil, err
		}

		users = append(users, friend)
	}

	return users, err
}

func (r *Repository) AddFriend(userId int, fiendId int) error {
	query := `insert into user_friend values(?, ?, ?), (?, ?, ?)`
	_, err := r.db.ExecContext(r.createContext(), query, userId, fiendId, "friend", fiendId, userId, "friend")

	return err
}

func (r *Repository) RemoveFriend(userId int, friendId int) error {
	query := `delete from user_friend where userFrom=? and userTo=?`
	_, err := r.db.ExecContext(r.createContext(), query, userId, friendId)

	if err != nil {
		return err
	}

	query = `delete from user_friend where userFrom=? and userTo=?`
	_, err = r.db.ExecContext(r.createContext(), query, friendId, userId)

	return err
}

func (r *Repository) GetFriendList(userId int, skip int, take int) ([]user.Friend, int, error) {
	query := `select user.id, user.name, user.surname, user.avatar from user left join user_friend on userFrom=user.id 
				where userTo=? LIMIT ?, ?`

	rows, err := r.db.QueryContext(r.createContext(), query, userId, skip, take)

	if err != nil {
		return nil, 0, err
	}

	friendList := make([]user.Friend, 0)

	for rows.Next() {
		var friend user.Friend

		if err = scanFriend(rows, &friend); err != nil {
			return nil, 0, err
		}

		friendList = append(friendList, friend)
	}
	queryTotal := `select count(userFrom) from user_friend where userFrom=?`

	row := r.db.QueryRowContext(r.createContext(), queryTotal, userId)
	var total int
	err = row.Scan(&total)

	return friendList, total, err
}

func (r *Repository) fullUserQuery(form user.FindUsersForm) (*sql.Rows, error) {
	query := selectStr + " where name like concat(?, '%') and surname like concat(?, '%')"
	query = addOrder(query)
	rows, err := r.db.QueryContext(r.createContext(), query, form.Name, form.Surname, form.Skip, form.Take)

	return rows, err
}

func (r *Repository) findUserByQuery(form user.FindUsersForm, by string) (*sql.Rows, error) {
	query := fmt.Sprintf("%s where %s like concat(?, '%%')", selectStr, by)
	query = addOrder(query)
	if by == "name" {
		return r.db.QueryContext(r.createContext(), query, form.Name, form.Skip, form.Take)
	}

	return r.db.QueryContext(r.createContext(), query, form.Surname, form.Skip, form.Take)
}

func (r *Repository) findUserQuery(form user.FindUsersForm) (*sql.Rows, error) {
	query := addOrder(selectStr)
	return r.db.QueryContext(r.createContext(), query, form.Skip, form.Take)
}

func scanFriend(rows *sql.Rows, friend *user.Friend) error {
	if err := rows.Scan(&friend.Id, &friend.Name, &friend.Surname, &friend.Avatar); err != nil {
		return err
	}

	return nil
}

func addOrder(query string) string {
	return query + " order by id DESC limit ?, ?"
}