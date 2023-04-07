package usecase

import (
	"github.com/NAVVG/test/internal/usecase/repo"
	"github.com/NAVVG/test/pkg/logger"
)

type Achivement struct {
	log  *logger.Logger
	repo repo.Achivement
}

func NewAchivement(log *logger.Logger, repo repo.Achivement) *Achivement {
	return &Achivement{
		log:  log,
		repo: repo,
	}
}

func (a *Achivement) ShowAchivement() {

}
