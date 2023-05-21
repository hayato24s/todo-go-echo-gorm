package middleware

import (
	"github.com/hayato24s/todo-echo-gorm/handler/common"
	"github.com/labstack/echo/v4"
)

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}
		code, ok := common.ErrToCode[err]
		if ok {
			return echo.NewHTTPError(code, err.Error())
		}
		return err
	}
}
