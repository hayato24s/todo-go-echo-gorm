package auth

import (
	"net/http"

	"github.com/hayato24s/todo-echo-gorm/handler/common"
	"github.com/labstack/echo/v4"
)

type LogOutRes struct{}

// LogOut
//
//	@Tags		logout
//	@Produce	json
//	@Success	200	{object}	LogOutRes
//	@Failure	401	{object}	common.ErrorRes
//	@Failure	500	{object}	common.ErrorRes
//	@Router		/logout [delete]
func (h *Handler) LogOut(c echo.Context) error {
	_, err := h.uc.Authenticate(c.Request().Context())
	if err != nil {
		return err
	}

	common.ClearUserIDInCookie(c)
	c.JSON(http.StatusOK, LogOutRes{})
	return nil
}
