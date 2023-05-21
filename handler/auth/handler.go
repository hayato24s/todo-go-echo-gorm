package auth

import (
	"github.com/hayato24s/todo-echo-gorm/usecase"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	uc *usecase.UseCase
}

func (h *Handler) Register(e *echo.Echo) {
	e.POST("/login", h.LogIn)
	e.DELETE("/logout", h.LogOut)
}

func NewHandler(uc *usecase.UseCase) *Handler {
	return &Handler{uc: uc}
}
