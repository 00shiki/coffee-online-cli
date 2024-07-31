package users

import "coffee-online-cli/entity"

type Repo interface {
	Reader
	Writer
}

type Reader interface {
	CheckEmailExists(email string) error
	GetUserByEmail(email string) (*entity.User, error)
	GetUserByID(id int) (*entity.User, error)
}

type Writer interface {
	CreateUser(user entity.User) error
}
