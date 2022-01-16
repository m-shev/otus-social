package post_storage

import (
	"context"
	"database/sql"
	"github.com/m-shev/otus-social/back/internal/connector"
	"github.com/m-shev/otus-social/back/internal/services/post"
	"strings"
	"time"
)

const queryTimeout = time.Second * 15

type Storage struct {
	db *sql.DB
}

func NewRepository(con *connector.Connector) *Storage {
	return &Storage{db: con.GetConnection()}
}

func (s *Storage) Create(form post.CreatePostForm) (post.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()

	query := `insert into post(authorId, content, imageLink) values(?, ?, ?)`

	result, err := s.db.ExecContext(ctx, query, form.AuthorId, form.Content, form.ImageLink)

	if err != nil {
		return post.Post{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return post.Post{}, err
	}

	return s.GetById(int(id))
}

func (s *Storage) GetById(id int) (post.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	query := `select * from post where id=?`

	row := s.db.QueryRowContext(ctx, query, id)

	p := post.Post{}

	err := row.Scan(&p.Id, &p.AuthorId, &p.Content, &p.ImageLink, &p.CreateAt, &p.UpdateAt)

	return p, err
}

func (s *Storage) GetList(ids []int) ([]post.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	l := len(ids)

	postList := make([]post.Post, 0)

	query := "select * from post where id in (?" + strings.Repeat(",?", l-1) + ") order by createAt desc"
	args := make([]interface{}, 0, l)

	for _, v := range ids {
		args = append(args, v)
	}

	rows, err := s.db.QueryContext(ctx, query, args...)

	if err != nil {
		return postList, err
	}

	for rows.Next() {
		var p post.Post

		if err = rows.Scan(&p.Id, &p.AuthorId, &p.Content, &p.ImageLink, &p.CreateAt, &p.UpdateAt); err != nil {
			return postList, err
		}

		postList = append(postList, p)
	}

	return postList, nil
}
