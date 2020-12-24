package components

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Password struct{ Password string }

func (p Password) HashAndSalt() string {
	hash, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func (p Password) ComparePasswords(plainPwd []byte) bool {
	byteHash := []byte(p.Password)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
