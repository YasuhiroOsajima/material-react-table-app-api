package model

type User struct {
	Username string `json:"usernames"`
	Age      int    `json:"age"`
}

func NewUser(username string, age int) *User {
	return &User{Username: username, Age: age}
}
