package server

import "github.com/gojwt/models"

type ServerI interface {
	Init() (err error)
	CreateUser(*models.User) (err error)
}
