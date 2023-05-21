package tasks

import (
	"net/http"

	"github.com/hayato24s/todo-echo-gorm/appctx"
	"github.com/hayato24s/todo-echo-gorm/apperr"
	"github.com/hayato24s/todo-echo-gorm/usecase"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UpdateBody struct {
	Title     *string `json:"title" example:"read documentation" extensions:"x-nullable"`
	Completed *bool   `json:"completed" example:"false" extensions:"x-nullable"`
}

type UpdateRes struct{}

// Update
//
//	@Tags		tasks
//	@Accept		json
//	@Produce	json
//	@Param		id	path		uint		true	"Task ID"
//	@Param		in	body		UpdateBody	true	"request body"
//	@Success	200	{object}	UpdateRes
//	@Failure	400	{object}	common.ErrorRes
//	@Failure	401	{object}	common.ErrorRes
//	@Failure	404	{object}	common.ErrorRes
//	@Failure	500	{object}	common.ErrorRes
//	@Router		/tasks/{id} [put]
func (h *Handler) Update(c echo.Context) error {
	_, ok := appctx.GetUserID(c.Request().Context())
	if !ok {
		return apperr.ErrUnauthorized
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var in *UpdateBody
	err = c.Bind(&in)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = h.uc.UpdateTask(c.Request().Context(), usecase.UpdateTaskIn{ID: id, Title: in.Title, Completed: in.Completed})
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, UpdateRes{})
	return nil
}
