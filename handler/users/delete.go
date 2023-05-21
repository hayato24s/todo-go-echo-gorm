package users

import (
	"net/http"

	"github.com/hayato24s/todo-echo-gorm/handler/common"
	"github.com/labstack/echo/v4"
)

type DeleteRes struct{}

// Delete
//
//	@Tags		users
//	@Produce	json
//	@Success	200	{object}	DeleteRes
//	@Failure	401	{object}	common.ErrorRes
//	@Failure	500	{object}	common.ErrorRes
//	@Router		/users/me [delete]
func (h *Handler) Delete(c echo.Context) error {
	err := h.uc.DeleteUser(c.Request().Context())
	if err != nil {
		return err
	}

	common.ClearUserIDInCookie(c)
	c.JSON(http.StatusOK, DeleteRes{})
	return nil
}
