package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Api struct {
	Router *mux.Router
}

func (self *Api) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>HomePage</h1>")
	fmt.Println("Endpoint Hit: homePage")
}

func (self *Api) Serve() {
	self.Router = mux.NewRouter().StrictSlash(true)
	self.Router.HandleFunc("/", self.homePage)
	http.HandleFunc("/", self.homePage)
	log.Fatal(http.ListenAndServe(":3000", self.Router))
}
