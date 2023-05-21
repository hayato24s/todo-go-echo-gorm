package tasks

import (
	"net/http"

	"github.com/hayato24s/todo-echo-gorm/appctx"
	"github.com/hayato24s/todo-echo-gorm/apperr"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type DeleteRes struct{}

// Delete
//
//	@Tags		tasks
//	@Produce	json
//	@Param		id	path		uint	true	"Task ID"
//	@Success	200	{object}	DeleteRes
//	@Failure	401	{object}	common.ErrorRes
//	@Failure	404	{object}	common.ErrorRes
//	@Failure	500	{object}	common.ErrorRes
//	@Router		/tasks/{id} [delete]
func (h *Handler) Delete(c echo.Context) error {
	_, ok := appctx.GetUserID(c.Request().Context())
	if !ok {
		return apperr.ErrUnauthorized
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = h.uc.DeleteTask(c.Request().Context(), id)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, DeleteRes{})
	return nil
}
