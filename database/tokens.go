package database

import (
	"context"
	"fmt"
	"os"

	"github.com/gojwt/models"
)

func (s *Postgres) SaveTokenForUser(usr models.User, token string) (err error) {
	conn, err := s.Pool.Acquire(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), "insert into tokens  (creationdate,expirationdate,TOKEN,userid) values(current_timestamp,NULL,$1,$2)", token, usr.Id)

	return err
}
