package database

import "github.com/gojwt/models"

/* Primary Data access contract to be fullfilled*/
type PostgresI interface {
	Init() (err error)
	CreateUser(*models.User) (err error)
	GetUserById(id int) (usr models.User, err error)
	GetUserByUserName(username string) (usr models.User, err error)
	SaveTokenForUser(usr models.User, token string) error
}
