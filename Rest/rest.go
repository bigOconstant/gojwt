package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gojwt/components"
	"github.com/gojwt/database"
	"github.com/gojwt/models"
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
	Message string `json:"message"`
	Token   string `json:"token"`
}

type CreateUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
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
	response := LoginResponse{Success: false, Message: "", Token: ""}

	fmt.Println("calling login")
	var i LoginIn
	json.NewDecoder(r.Body).Decode(&i)

	usr, err := self.DB.GetUserByUserName(i.Username)
	if err != nil {
		response.Message = "error getting user"
		fmt.Println("error getting user")
		msg, _ := json.Marshal(response)
		fmt.Fprintf(w, string(msg))
		return
	}

	matches := components.Password{Password: usr.Password}.ComparePasswords([]byte(i.Password))

	if matches {
		fmt.Println("Password matches login user")
		response.Message = "Password matches"
		tokenC, err := models.CreateTokenClaims(usr)
		if err == nil {
			response.Token = tokenC
			response.Success = true
			response.Message = "Token Creation Success"
			err := self.DB.SaveTokenForUser(usr, tokenC)
			if err != nil {
				response.Success = false
				response.Message = "Token Creation Failed"
				response.Token = ""
			}

		} else {
			response.Message = "error creating token"
			fmt.Println(err)
		}

	} else {
		response.Message = "password doesn't match"
	}

	msg, _ := json.Marshal(response)
	fmt.Fprintf(w, string(msg))

}

/*
CreateUser: steps,
1.Get username and password from request body.
2. check to make sure username doesn't already exist
3. if Create user send message saing create user operation was a success
*/
func (self *Api) createUser(w http.ResponseWriter, r *http.Request) {
	response := CreateUserResponse{Success: false, Message: ""}
	fmt.Println("Calling create user")
	var i CreateUser
	json.NewDecoder(r.Body).Decode(&i)

	userExist := self.DB.UserExist(i.Username)
	if userExist {
		fmt.Println("user exist")
		response.Message = "user exist"

	} else {
		fmt.Println("User doesn't exist. Creating user")
		pw := components.Password{Password: i.Password}.HashAndSalt()
		var User models.User = models.User{Username: i.Username,
			Password: pw,
			Email:    i.Email,
			Data:     i.Data,
		}
		err := self.DB.CreateUser(&User)
		if err == nil {
			response.Message = "User Created"
			response.Success = true
		} else {
			response.Message = "Error Creating user"
			fmt.Println(err)
		}

	}
	msg, _ := json.Marshal(response)
	fmt.Fprintf(w, string(msg))

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
