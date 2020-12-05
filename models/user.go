package models

type User struct {
	Id       string      `json:"id"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Data     interface{} `json:"data"` //Unstructured jsonb data. What ever someone wants to store.
}
