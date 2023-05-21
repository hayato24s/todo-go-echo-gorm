package usecase

import (
	"github.com/hayato24s/todo-echo-gorm/port"
)

type UseCase struct {
	r port.IRepository
}

func NewUseCase(r port.IRepository) *UseCase {
	return &UseCase{r: r}
}
