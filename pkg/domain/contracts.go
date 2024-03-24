package domain

type Store interface {
	CreateUser(usr *User) (*User, error)
	DeleteUserByID(id int64) error
	FindUserByEmail(email string) (*User, error)
	FindUserByID(id int64) (*User, error)
}
