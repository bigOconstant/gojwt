package components

import (
	"testing"

	"github.com/gojwt/models"
)

func TestComparePasswords(t *testing.T) {
	var m map[string]string
	m = make(map[string]string)
	m["firstname"] = "nightWing"
	m["lastname"] = "duck"
	var usr models.User = models.User{Id: 0, Username: "duckwing", Password: Password{}.HashAndSalt([]byte("testpassword")), Email: "duck@ducksauce.com", Data: m}

	compared := Password{}.ComparePasswords(usr.Password, []byte("testpassword"))
	if compared {
		t.Log("True")
	} else {
		t.Error("Doesn't work")
	}

}
