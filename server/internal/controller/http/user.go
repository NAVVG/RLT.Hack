package http

import "github.com/NAVVG/test/internal/usecase"

type User struct {
	usecase usecase.User
}

func NewUser(usecase usecase.User) *User {
	return &User{
		usecase: usecase,
	}
}
