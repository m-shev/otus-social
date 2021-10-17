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
	Id   int
	Name string
	Surname   string
	Age       uint8
	City      string
	Email     string
	Password  string
	Interests []Interest
}

type Interest struct {
	Id   int
	Name string
}
