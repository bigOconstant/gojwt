package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Api struct {
	Router *mux.Router
}

type LoginIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (self *Api) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>HomePage</h1>")
	fmt.Println("Endpoint Hit: homePage")
}

//todo. Ch
func (self *Api) login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("calling login")
	var i LoginIn
	json.NewDecoder(r.Body).Decode(&i)
	fmt.Fprintf(w, "username: %s ,password:%s\n", i.Username, i.Password)

}

func (self *Api) Serve() {
	self.Router = mux.NewRouter().StrictSlash(true)
	self.Router.HandleFunc("/", self.homePage)
	http.HandleFunc("/", self.homePage)
	self.Router.HandleFunc("/login", self.login).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", self.Router))
}

//r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
