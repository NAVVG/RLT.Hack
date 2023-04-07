package repo

import (
	"github.com/NAVVG/test/internal/entity"
	"github.com/NAVVG/test/pkg/postgres"
)

type Аchievement struct {
	postgres *postgres.Postgres
}

func NewАchievement(pg *postgres.Postgres) *Аchievement {
	return &Аchievement{postgres: pg}
}

func (a *Аchievement) CreateDescription(description string) error

func (a *Аchievement) CreateIcon(icon []byte) error

func (a *Аchievement) CreateHeader(header string) error

func (a *Аchievement) GetById() entity.Аchievement
