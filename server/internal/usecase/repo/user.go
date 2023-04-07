package repo

import (
	"github.com/NAVVG/test/internal/entity"
	"github.com/NAVVG/test/pkg/postgres"
)

type User struct {
	postgres *postgres.Postgres
}

func (u *User) Create(name string, email string, password string) error

func (u *User) GetByName(name string) (entity.User, error)

func (u *User) DeleteByName(name string) error

func NewUser(pg *postgres.Postgres) *User {
	return &User{postgres: pg}
}
