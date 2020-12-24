package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gojwt/database"
	"github.com/gorilla/mux"
)

type Api struct {
	Router *mux.Router
	DB     database.PostgresI
}

type LoginIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUser struct {
	Username string      `json:"username"`
	Password string      `json:"password"`
	Email    string      `json:"email"`
	Data     interface{} `json:"data"`
}

/*
CREATE TABLE IF NOT EXISTS users (
 id SERIAL NOT NULL PRIMARY KEY,
 CreateDate timestamp,
 username varchar(160) NOT NULL UNIQUE,
 password text NOT NULL,
 email varchar(160) NOT NULL,
 data jsonb)

*/

type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

func (self *Api) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>HomePage</h1>")
	fmt.Println("Endpoint Hit: homePage")
}

/*
login: steps
1. Call get user by username
2. Chech password, by calling comparepassword method.
3. If pasword matches, create token claims,
4. return stringified token claims aka jwt

*/
func (self *Api) login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("calling login")
	var i LoginIn
	json.NewDecoder(r.Body).Decode(&i)

	usr, err := self.DB.GetUserByUserName(i.Username)
	if err != nil {
		fmt.Println("error getting user")
	}

	fmt.Println("userid:", usr.Id)
	fmt.Fprintf(w, "username: %s ,password:%s\n", i.Username, i.Password)

}

/*
CreateUser: steps,
1.Get username and password from request body.
2. check to make sure username doesn't already exist
3. if Create user send message saing create user operation was a success
*/
func (self *Api) createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Calling create user")
	var i CreateUser
	json.NewDecoder(r.Body).Decode(&i)
	fmt.Println("user's email:", i.Email)
	usr, err := self.DB.GetUserByUserName(i.Username)
	if err != nil && err.Error() == "no rows in result set" {

		fmt.Println("user does not exist")

	}
	fmt.Println(usr.Id)

}

func (self *Api) Serve() {
	self.DB.Init()
	self.Router = mux.NewRouter().StrictSlash(true)
	self.Router.HandleFunc("/", self.homePage)
	http.HandleFunc("/", self.homePage)
	self.Router.HandleFunc("/login", self.login).Methods("POST")
	self.Router.HandleFunc("/CreateUser", self.createUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", self.Router))
}

//r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
