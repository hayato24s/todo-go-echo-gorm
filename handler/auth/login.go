package auth

import (
	"net/http"

	"github.com/hayato24s/todo-echo-gorm/handler/common"
	"github.com/hayato24s/todo-echo-gorm/usecase"
	"github.com/labstack/echo/v4"
)

type LogInBody struct {
	Name     string `json:"name" example:"gopher"`
	Password string `json:"password" minLength:"8" maxLength:"20" example:"password"`
}

type LogInRes struct{}

// LogIn
//
//	@Tags		login
//	@Accept		json
//	@Produce	json
//	@Param		body	body		LogInBody	true	"request body"
//	@Success	200		{object}	LogInRes
//	@Failure	400		{object}	common.ErrorRes
//	@Failure	500		{object}	common.ErrorRes
//	@Router		/login [post]
func (h *Handler) LogIn(c echo.Context) error {
	var body LogInBody
	err := c.Bind(&body)
	if err != nil {
		return err
	}

	user, err := h.uc.LogIn(c.Request().Context(), &usecase.LogInIn{Name: body.Name, Password: body.Password})
	if err != nil {
		return err
	}

	err = common.SetUserIDInCookie(c, user.ID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, LogInRes{})
	return nil
}
