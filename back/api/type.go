package api

import "github.com/m-shev/otus-social/back/internal/services/user"

type friendListResponse struct {
	FriendList []user.Friend `json:"friendList"`
	Total      int           `json:"total"`
}

type userListResponse struct {
	UserList []user.Friend `json:"userList"`
	Total    int           `json:"total"`
}
