package usecase

import (
	"github.com/NAVVG/RLT.Hack/internal/usecase/repo"
)

type User struct {
	repo repo.User
}

func NewUser(repo repo.User) *User {
	return &User{
		repo: repo,
	}
}
