package tasks

import (
	"time"

	"github.com/hayato24s/todo-echo-gorm/entity"
)

type TaskRes struct {
	ID        string    `json:"id" example:"de0bf6f0-a09a-4e0d-aaa4-b1bf4d953d1e"`
	Title     string    `json:"title" example:"read documentation"`
	Completed bool      `json:"completed" example:"false"`
	CreatedAt time.Time `json:"created_at" example:"2006-01-02T15:04:05Z"`
}

func ToTaskRes(t *entity.Task) TaskRes {
	return TaskRes{
		ID:        t.ID.String(),
		Title:     t.Title,
		Completed: t.Completed,
		CreatedAt: t.CreatedAt,
	}
}
