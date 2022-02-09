package post

import (
	"golang.org/x/net/context"
	"strconv"
)

func (s *Service) RecreateCache() {
	s.readCreatePost = false
	s.logger.Printf("stopping read queue create cache \n")

	<-s.finishedReading

	s.logger.Printf("start recreate cache")
	// do something
	s.readCreatePost = true
	go s.ConsumeCreateMessage()
}

func (s *Service) recreateCache() error {
	done := false
	skip := 0
	take := 10

	for !done {
		ids, total, err := s.readUsers(skip, take)

		if err != nil {
			return err
		}

		for _, id := range ids {

		}

		skip += take

		if total < skip {
			done = true
		}
	}

}

func (s *Service) fillUserCache(userId int) error {
	done := false
	skip := 0
	take := 50
	key := strconv.Itoa(userId)
	postIds, err := s.getConsumerCache(key)

	if err != nil {
		return nil
	}

	for !done {
		posts, total, err := s.readPosts(userId, skip, take)

		if err != nil {
			return err
		}

		for _, post := range posts {
			contentCache, err := s.getContentCache(post)

			if err != nil {
				return err
			}

			postIds = append(postIds, post.Id)
		}

		skip += take

		if total < skip {
			done = true
		}
	}
}

func (s *Service) readPosts(userId, take, skip int) ([]Content, int, error) {
	postList := make([]Content, 0, take)

	query := `select p.id, p.authorId, p.content, p.imageLink, p.createAt, p.updateAt from user
              join user_friend uf on user.id = uf.userTo
              join post p on uf.userFrom = p.authorId
		      where user.id=? order by p.createAt DESC LIMIT 1000`

	totalQuery := `select count(p.id) from user
              join user_friend uf on user.id = uf.userTo
              join post p on uf.userFrom = p.authorId
		      where user.id=? order by p.createAt DESC LIMIT 1000`

	rows, err := s.db.QueryContext(context.Background(), query, userId, skip, take)

	if err != nil {
		return postList, 0, err
	}

	for rows.Next() {
		var post Content
		err = rows.Scan(&post.Id, &post.AuthorId, &post.Content, &post.CreateAt, &post.UpdateAt)

		if err != nil {
			return postList, 0, err
		}

		postList = append(postList, post)
	}

	var total int

	row := s.db.QueryRowContext(context.Background(), totalQuery)

	err = row.Scan(&total)

	return postList, total, err
}

func (s *Service) readUsers(skip, take int) ([]int, int, error) {
	query := `select user.id, count(userTo) as friends  from user
				left join user_friend uf on user.id = uf.userTo
				group by user.id order by friends DESC LIMIT ?,?`

	totalQuery := `select count(user.id) from user`

	rows, err := s.db.QueryContext(context.Background(), query, skip, take)

	if err != nil {
		return []int{}, 0, err
	}

	ids := make([]int, 0, take)

	for rows.Next() {
		var id int
		err = rows.Scan(&id)

		if err != nil {
			return nil, 0, err
		}

		ids = append(ids, id)
	}

	var total int

	row := s.db.QueryRowContext(context.Background(), totalQuery)

	err = row.Scan(&total)

	return ids, total, err
}
