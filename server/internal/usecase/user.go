package usecase

import (
	"github.com/NAVVG/test/internal/usecase/repo"
)

type User struct {
	repo repo.User
}

func NewUser(repo repo.User) *User {
	return &User{
		repo: repo,
	}
}
