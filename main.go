package main

import (
	rest "github.com/gojwt/Rest"
)

func main() {
	// fmt.Println("hello world")
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"foo": "bar",
	// 	"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	// })
	// tokenString, err := token.SignedString([]byte("ducksauce"))
	// fmt.Println(tokenString, err)
	// var S = database.Postgres{}
	// var i database.PostgresI = &S
	// err = i.Init()
	// if err != nil {
	// 	println(err.Error())
	// }
	// var m map[string]string
	// m = make(map[string]string)
	// m["firstname"] = "caleb"
	// m["lastname"] = "mccarthy"
	// var usr models.User = models.User{Id: 0, Username: "andrew", Password: components.Password{}.HashAndSalt([]byte("testpassword")), Email: "duck@ducksauce.com", Data: m}

	// err = i.CreateUser(&usr)
	// if err != nil {
	// 	println(err.Error())
	// }

	// returnedUser, err := i.GetUserByUserName("andrew")

	// fmt.Println("compared:", components.Password{}.ComparePasswords(returnedUser.Password, []byte("testpassword")))

	// fmt.Println(string(returnedUser.DataToJson()))

	// jwtstring, err := models.CreateTokenClaims(returnedUser)
	// fmt.Println("jwtstring:", jwtstring)

	// if err != nil {
	// 	fmt.Println("error here", err.Error())

	// } else {
	// 	models.CreateClaimFromTokenString(jwtstring)

	// 	i.SaveTokenForUser(returnedUser, jwtstring)
	// }

	Server := rest.Api{}
	Server.Serve()

}
