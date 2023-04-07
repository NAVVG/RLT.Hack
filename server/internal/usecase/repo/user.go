package repo

import (
	"context"

	"github.com/NAVVG/test/internal/entity"
	"github.com/NAVVG/test/pkg/postgres"
)

type User struct {
	postgres *postgres.Postgres
}

func (u *User) Create(name string, email string, password string) error {
	query := "insert into users (name, email, password) values ($1, $2, $3)"
	_, err := u.postgres.Pool.Exec(context.Background(), query, name, email, password)
	return err
}

func (u *User) GetById(id int) (entity.User, error) {
	query := "select * from users where id=$1"
	var user entity.User
	err := u.postgres.Pool.QueryRow(context.Background(), query, id).Scan(&user)
	return user, err
}

func (u *User) DeleteById(id int) error {
	query := "delete from users where id=$1"
	_, err := u.postgres.Pool.Exec(context.Background(), query, id)
	return err
}

func NewUser(pg *postgres.Postgres) *User {
	return &User{postgres: pg}
}
