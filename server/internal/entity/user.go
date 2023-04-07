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

type Token struct {
	User    int
	Token   string
	ExpDate time.Time
}
