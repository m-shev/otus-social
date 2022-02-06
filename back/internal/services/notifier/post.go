package notifier

import (
	"fmt"
	"github.com/m-shev/otus-social/back/internal/services/post"
)

func (s *Service) PostCreated(authorId int, p post.Post) {
	done := false
	skip := 0

	for !done {
		friends, total, err := s.userService.GetFriendList(authorId, takeFriends, skip)

		if err != nil {
			s.logger.Println(s.sendPostCreatedError(authorId, p.Id, err))
			done = true
		}

		consumerIds := make([]int, 0, len(friends))

		for _, v := range friends {
			consumerIds = append(consumerIds, v.Id)
		}

		err = s.postQueue.SendPostCreated(MessagePostCreate{
			AuthorId:  authorId,
			Post:      p,
			Consumers: consumerIds,
		})

		if err != nil {
			s.logger.Println(s.sendPostCreatedError(authorId, p.Id, err))
			done = true
		}

		skip += takeFriends

		if total < skip {
			done = true
		}
	}
}

func (s *Service) sendPostCreatedError(authorId int, postId int, err error) string {
	return fmt.Sprintf("warn: cant send message post %d from user %d has been created, cause %s",
		postId, authorId, err.Error())
}
