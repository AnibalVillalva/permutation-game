package usecases

import (
	"context"

	log "github.com/sirupsen/logrus"

	"permutation-game/api/cmd/internal/repositories"
)

type UseCase struct {
	repo repositories.Getter
}

func New(r repositories.Getter) *UseCase {
	return &UseCase{r}
}
func (uc UseCase) Execute(ctx context.Context) (context.Context, error) {
	log.Debugf("[UseCase] Starting.")
	return uc.repo.Get(ctx)
}
