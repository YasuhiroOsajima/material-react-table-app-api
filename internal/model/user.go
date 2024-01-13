package model

type User struct {
	Username string `json:"name"`
	Age      int    `json:"age"`
}

func NewUser(username string, age int) *User {
	return &User{Username: username, Age: age}
}
