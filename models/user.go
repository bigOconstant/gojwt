package models

import "encoding/json"

type User struct {
	Id       int         `json:"id"`
	Username string      `json:"username"`
	Password string      `json:"password,omitempty"`
	Email    string      `json:"email"`
	Data     interface{} `json:"data,omitempty"` //Unstructured jsonb data. What ever someone wants to store.
}

func (u *User) DataToJson() []byte {
	val, _ := json.Marshal(u.Data)
	return val
}
