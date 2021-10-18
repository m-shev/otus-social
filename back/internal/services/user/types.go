package user

type Repository interface {
	Create(form CreateUserForm) (User, error)
	GetById(id int) (User, error)
	FindUser(email string) (User, error)
}

type CreateUserForm struct {
	Name      string
	Surname   string
	Age       uint8
	City      string
	Email     string
	Password  string
	Interests []string
}

type User struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Surname   string     `json:"surname"`
	Age       uint8      `json:"age"`
	City      string     `json:"city"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Interests []Interest `json:"interests"`
}

type Interest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
