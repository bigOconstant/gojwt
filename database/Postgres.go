package database

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

//Postgres Implementation of ServerI Data Access Layer

type Postgres struct {
	Host     string
	Database string
	UserName string
	Password string
	Port     string

	/* Most things can be done with a thread safe connection pool
	after the initial connection and i */
	Pool *pgxpool.Pool
}

//GetConnection Gets connection to postgres
func (s Postgres) GetConnection() (conn *pgx.Conn, err error) {
	databaseurl := "postgresql://" + s.UserName + ":" + s.Password + "@" + s.Host + ":" + s.Port + "/" + s.Database
	return pgx.Connect(context.Background(), databaseurl)
}

//getConnectionNoDatabase Gets connection to postgres without specifying database
func (s Postgres) getConnectionNoDatabase() (conn *pgx.Conn, err error) {
	databaseurl := "postgresql://" + s.UserName + ":" + s.Password + "@" + s.Host + ":" + s.Port
	return pgx.Connect(context.Background(), databaseurl)
}

func (self *Postgres) Init() (err error) {
	self.Password = os.Getenv("PGPASSWORD")
	self.UserName = os.Getenv("PGUSERNAME")
	self.Host = os.Getenv("PGHOST")
	self.Database = os.Getenv("PGDB")
	self.Port = os.Getenv("PGPORT")
	fmt.Println("password:", self.Password)
	self.initDataBase()
	databaseurl := "postgresql://" + self.UserName + ":" + self.Password + "@" + self.Host + ":" + self.Port + "/" + self.Database
	self.Pool, err = pgxpool.Connect(context.Background(), databaseurl)
	if err == nil {
		self.initTables()
	}
	return err
}

//initDataBase Responsible for initializing Database
func (s *Postgres) initDataBase() {
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
		// Don't use connection pool, it's not set up yet
		conn, err := s.getConnectionNoDatabase()
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

func (s *Postgres) initTables() {
	tableNames := s.getTables()
	if val, ok := tableNames["users"]; ok {
		fmt.Println("users table found", val)
	} else {
		fmt.Println("users table not found")
		query, err := ioutil.ReadFile("sqlFiles/CreateUserTable.sql")
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
			fmt.Println("Error creating User table")
		} else {
			fmt.Println("Success creating User table")
		}

	}

	if val, ok := tableNames["tokens"]; ok {
		fmt.Println("tokens table found", val)
	} else {
		fmt.Println("tokens table not found")
		query, err := ioutil.ReadFile("sqlFiles/CreateTokenTable.sql")
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

			fmt.Println("Error creating token table")
			fmt.Println(err)
		} else {
			fmt.Println("Success creating token table")
		}

	}
}

func (s *Postgres) getDatabases() map[string]bool {
	var names map[string]bool = make(map[string]bool)
	// Don't use connection pool, because we havn't set it up yet
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

func (s *Postgres) getTables() map[string]bool {
	var names map[string]bool = make(map[string]bool)
	conn, err := s.Pool.Acquire(context.Background())
	//conn, err := s.getConnectionNoDatabase()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return names
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), "SELECT tablename FROM pg_catalog.pg_tables;")

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
		} else {
			fmt.Println("name:", name)
		}
		names[name] = true

	}

	return names
}
