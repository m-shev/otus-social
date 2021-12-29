package dialog

import "time"

type Dialog struct {
	DialogId  int64
	Name      string
	CreatorId int
	CreateAt  time.Time
}

type CreateDialogForm struct {
	Name   string `json:"name"`
	UserId int    `json:"userId"`
}

type Member struct {
	MemberId int
	DialogId int64
	Role     string
	CreateAt time.Time
}

type AddMemberForm struct {
	DialogId int64  `json:"dialogId"`
	MemberId int    `json:"memberId"`
	Role     string `json:"role"`
}
