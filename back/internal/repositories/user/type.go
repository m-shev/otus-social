package user

type CreateUserForm struct {
	Name string
	Surname string
	Age uint8
	City string
	Email string
	Password string
	Interests []string
}

type User struct {
	id int
	Name string
	Surname string
	Age uint8
	City string
	Email string
	Password string
	Interests []string
}

type Interest struct {
	Id int
	Name string
}