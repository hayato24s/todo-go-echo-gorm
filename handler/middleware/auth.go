package middleware

import (
	"github.com/hayato24s/todo-echo-gorm/appctx"
	"github.com/hayato24s/todo-echo-gorm/handler/common"
	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := common.GetUserIDFromCookie(c)
		if err == nil {
			common.SetUserIDInCookie(c, userID)
			c.SetRequest(c.Request().Clone(appctx.SetUserID(c.Request().Context(), userID)))
		}
		return next(c)
	}
}
