package main

import (
	rest "github.com/gojwt/Rest"
	"github.com/gojwt/database"
)

func main() {

	var S = database.Postgres{}

	Server := rest.Api{DB: &S}
	Server.Serve()

}
