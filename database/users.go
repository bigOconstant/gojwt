package database

import (
	"context"
	"fmt"
	"os"

	"github.com/gojwt/models"
)

func (s *Postgres) CreateUser(usr *models.User) (err error) {
	conn, err := s.Pool.Acquire(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), "insert into users  (username,password, email, data,createdate) values($1,$2,$3,$4,current_timestamp)", usr.Username, usr.Password, usr.Email, usr.DataToJson())

	return err
}

func (s *Postgres) GetUserById(id int) (usr models.User, err error) {
	conn, err := s.Pool.Acquire(context.Background())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return usr, err
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(), "SELECT id,username,password,email,data FROM users where id = $1", id)
	err = row.Scan(&usr.Id, &usr.Username, &usr.Password, &usr.Email, &usr.Data)
	if err != nil {
		fmt.Println(err.Error())
	}
	return usr, err
}
func (s *Postgres) GetUserByUserName(username string) (usr models.User, err error) {
	conn, err := s.Pool.Acquire(context.Background())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return usr, err
	}
	defer conn.Release()

	row := conn.QueryRow(context.Background(), "SELECT id,username,password,email,data FROM users where username = $1", username)
	err = row.Scan(&usr.Id, &usr.Username, &usr.Password, &usr.Email, &usr.Data)

	if err != nil {
		fmt.Println(err.Error())
	}
	return usr, err
}
