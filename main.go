package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gojwt/models"
	"github.com/gojwt/server"
	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
)

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func main() {
	fmt.Println("hello world")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, err := token.SignedString([]byte("ducksauce"))
	fmt.Println(tokenString, err)
	var S = server.Postgres{}
	var i server.ServerI = &S
	err = i.Init()
	if err != nil {
		println(err.Error())
	}
	var m map[string]string
	m = make(map[string]string)
	m["firstname"] = "caleb"
	m["lastname"] = "mccarthy"
	var usr models.User = models.User{Id: 0, Username: "andrew", Password: hashAndSalt([]byte("testpassword")), Email: "duck@ducksauce.com", Data: m}

	err = i.CreateUser(&usr)
	if err != nil {
		println(err.Error())
	}

	returnedUser, err := i.GetUserByUserName("andrew")

	fmt.Println("compared:", comparePasswords(returnedUser.Password, []byte("testpassword")))

	fmt.Println(string(returnedUser.DataToJson()))

	jwtstring, err := models.CreateTokenClaims(returnedUser)
	fmt.Println("jwtstring:", jwtstring)

	if err != nil {
		fmt.Println("error here", err.Error())

	} else {
		models.CreateClaimFromTokenString(jwtstring)
	}

}
