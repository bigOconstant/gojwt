package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	pgx "github.com/jackc/pgx/v4"
)

//ServerI Interface to data Access Layer
type ServerI interface {
	NewServer() ServerI
	GetConnection() (conn *pgx.Conn, err error)
	InitDataBase()
}

//Server Implementation of Postgres Data Access Layer
type Server struct {
	Host     string
	Database string
	UserName string
	Password string
	Port     string
}

//GetConnection Gets connection to postgres
func (s Server) GetConnection() (conn *pgx.Conn, err error) {
	databaseurl := "postgresql://" + s.UserName + ":" + s.Password + "@" + s.Host + ":" + s.Port + "/" + s.Database
	return pgx.Connect(context.Background(), databaseurl)
}

//getConnectionNoDatabase Gets connection to postgres without specifying database
func (s Server) getConnectionNoDatabase() (conn *pgx.Conn, err error) {
	databaseurl := "postgresql://" + s.UserName + ":" + s.Password + "@" + s.Host + ":" + s.Port
	return pgx.Connect(context.Background(), databaseurl)
}

//InitDataBase Responsible for initializing Database
func (s Server) InitDataBase() {
	databases := s.getDatabases()

	if val, ok := databases["authdatabase"]; ok {
		fmt.Println("database found", val)
	} else {
		fmt.Println("Database not found")
		query, err := ioutil.ReadFile("sqlFiles/CreateDataBase.sql")
		if err != nil {
			fmt.Println("error reading in file")
			return
		}

		conn, err := s.GetConnection()
		if err != nil {
			fmt.Println("error getting connection")
			return
		}
		defer conn.Close(context.Background())

		_, err = conn.Exec(context.Background(), string(query))

		if err != nil {
			fmt.Println("Error creating database")
		} else {
			fmt.Println("Success creating database")
		}

	}

}

func (s Server) getDatabases() map[string]bool {
	var names map[string]bool = make(map[string]bool)
	conn, err := s.getConnectionNoDatabase()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return names
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT datname FROM pg_database")

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			fmt.Println("error")
			break
		}
		names[name] = true

	}

	return names
}

//NewServer Creates a new server object from env variables
func (s Server) NewServer() ServerI {
	var retVal Server = Server{}
	retVal.Password = os.Getenv("PGPASSWORD")
	retVal.UserName = os.Getenv("PGUSERNAME")
	retVal.Host = os.Getenv("PGHOST")
	retVal.Database = os.Getenv("PGDB")
	retVal.Port = os.Getenv("PGPORT")
	return retVal
}
