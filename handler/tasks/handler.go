package tasks

import (
	"github.com/hayato24s/todo-echo-gorm/usecase"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	uc *usecase.UseCase
}

func (h *Handler) Register(e *echo.Echo) {
	e.GET("/tasks", h.Index)
	e.POST("/tasks", h.Create)
	e.PUT("/tasks/:id", h.Update)
	e.DELETE("/tasks/:id", h.Delete)
}

func NewHandler(uc *usecase.UseCase) *Handler {
	return &Handler{uc: uc}
}
