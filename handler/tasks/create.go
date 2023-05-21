package tasks

import (
	"net/http"

	"github.com/hayato24s/todo-echo-gorm/usecase"
	"github.com/labstack/echo/v4"
)

type CreateBody struct {
	Title string `json:"title" example:"read documentation" validate:"required"`
}

type CreateRes TaskRes

// TasksCreate
//
//	@Tags		tasks
//	@Accept		json
//	@Produce	json
//	@Param		body	body		CreateBody	true	"request body"
//	@Success	200		{object}	CreateRes
//	@Failure	400		{object}	common.ErrorRes
//	@Failure	401		{object}	common.ErrorRes
//	@Failure	500		{object}	common.ErrorRes
//	@Router		/tasks [post]
func (h *Handler) Create(c echo.Context) error {
	var body CreateBody
	err := c.Bind(&body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	t, err := h.uc.CreateTask(c.Request().Context(), &usecase.CreateTaskIn{Title: body.Title})
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, ToTaskRes(t))
	return nil
}
