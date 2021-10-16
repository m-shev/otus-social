package user

type CreateUserForm struct {
	Name string
	Surname string
	Age uint8
	City string
	Email string
	Password string
}

type User struct {
	id int
	Name string
	Surname string
	Age uint8
	City string
	Email string
	Password string
}