package entity

import (
	"time"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

type Tokens struct {
	User    int
	Token   string
	ExpDate time
}
