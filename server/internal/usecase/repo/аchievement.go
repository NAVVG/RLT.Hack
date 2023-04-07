package repo

import (
	"github.com/NAVVG/RLT.Hack/pkg/postgres"
)

type Аchievement struct {
	postgres *postgres.Postgres
}

func NewАchievement(pg *postgres.Postgres) *Аchievement {
	return &Аchievement{postgres: pg}
}
