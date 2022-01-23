package user

type Repository interface {
	Create(form CreateUserForm) (User, error)
	GetById(id int) (User, error)
	FindUserByEmail(email string) (User, error)
	FindUsers(form FindUsersForm) ([]Friend, int, error)
	AddFriend(userId int, fiendId int) error
	RemoveFriend(userId int, fiendId int) error
	GetFriendList(userId int, skip int, take int) (friendList []Friend, total int, err error)
}

type CreateUserForm struct {
	Avatar    string
	Name      string
	Surname   string
	Age       uint8
	City      string
	Email     string
	Gender    string
	Password  string
	Interests []string
}

type AuthForm struct {
	Password string
	Login    string
}

type User struct {
	Profile
	Email    string `json:"email"`
	Password string `json:"-"`
}

type FriendForm struct {
	UserId   int `json:"userId"`
	FriendId int `json:"friendId"`
}

type FindUsersForm struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Skip    int    `json:"skip"`
	Take    int    `json:"take"`
}

type Friend struct {
	Id      int    `json:"id"`
	Avatar  string `json:"avatar"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Profile struct {
	Id        int        `json:"id"`
	Avatar    string     `json:"avatar"`
	Name      string     `json:"name"`
	Surname   string     `json:"surname"`
	Age       uint8      `json:"age"`
	Gender    string     `json:"gender"`
	City      string     `json:"city"`
	Interests []Interest `json:"interests"`
}

type Interest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
