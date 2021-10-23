package api

import "github.com/m-shev/otus-social/internal/services/user"

type friendListResponse struct {
	FriendList []user.Friend `json:"friendList"`
	Total      int           `json:"total"`
}
