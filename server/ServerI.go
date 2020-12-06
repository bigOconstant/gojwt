package server

import "github.com/gojwt/models"

type ServerI interface {
	Init() (err error)
	CreateUser(*models.User) (err error)
	GetUserById(id int) (usr models.User, err error)
	GetUserByUserName(username string) (usr models.User, err error)
}
