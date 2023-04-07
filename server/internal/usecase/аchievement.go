package usecase

import (
	"github.com/NAVVG/RLT.Hack/internal/usecase/repo"
	"github.com/NAVVG/RLT.Hack/pkg/logger"
)

type Аchievement struct {
	log  *logger.Logger
	repo repo.Аchievement
}

func NewАchievement(log *logger.Logger, repo repo.Аchievement) *Аchievement {
	return &Аchievement{
		log:  log,
		repo: repo,
	}
}

func (a *Аchievement) CreateAchivement(description, header string, icon []byte) {

}

func (a *Аchievement) ShowAchivement(id int) {

}
