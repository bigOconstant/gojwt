package models

import (
	"encoding/json"
	"time"
)

//User holds  the basic user object.
//Can stick other fields in Data interface
type User struct {
	Id         int         `json:"id"`
	Username   string      `json:"username"`
	Password   string      `json:"password,omitempty"`
	Email      string      `json:"email"`
	Data       interface{} `json:"data,omitempty"` //Unstructured jsonb data. What ever someone wants to store.
	CreateDate time.Time   `json:"createdate,omitempty"`
}

func (u *User) DataToJson() []byte {
	val, _ := json.Marshal(u.Data)
	return val
}
