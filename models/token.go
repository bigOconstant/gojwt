package models

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
	Id         int         `json:"id"`
	Username   string      `json:"username"`
	Email      string      `json:"email"`
	TokenCreateDate int64   `json:"tokencreatedate,omitempty"`
	Data       interface{} `json:"data,omitempty"` //Unstructured jsonb data. What ever someone wants to store.
	jwt.StandardClaims
}
func CreateTokenClaims(usr User) (signedToken string,err error) {
	claims := TokenClaims {
		Id:usr.Id,
		Username: usr.Username,
		Email: usr.Email,
		TokenCreateDate: time.Now().Unix(),
		Data:usr.Data,
		StandardClaims: jwt.StandardClaims{
			Issuer: "narnia",
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(1)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	signedToken, err = token.SignedString([]byte(os.Getenv("SECRETKEY")))
	return signedToken, err
}

func  CreateClaimFromTokenString(input string) (usr User, err error) {
	token, err := jwt.ParseWithClaims(
		input,
		&TokenClaims{},
		func(token *jwt.Token) (interface{},error) {
			return []byte(os.Getenv("SECRETKEY")),nil	
		},	
	)
	if err != nil {
		return User{},err
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return User{},errors.New("Couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return User{},errors.New("JWT is expired")
	}

	usr.Username = claims.Username
	usr.Email = claims.Email
	usr.Data = claims.Data

	return usr,nil
}
