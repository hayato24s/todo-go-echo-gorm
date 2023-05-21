package users

import (
	"net/http"

	"github.com/hayato24s/todo-echo-gorm/handler/common"
	"github.com/hayato24s/todo-echo-gorm/usecase"
	"github.com/labstack/echo/v4"
)

type CreateBody struct {
	Name     string `json:"name" example:"gopher"`
	Password string `json:"password" minLength:"8" maxLength:"20" example:"password"`
}

type CreateRes struct{}

// Create
//
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		body	body		CreateBody	true	"request body"
//	@Success	200		{object}	CreateRes
//	@Failure	400		{object}	common.ErrorRes
//	@Failure	500		{object}	common.ErrorRes
//	@Router		/users [post]
func (h *Handler) Create(c echo.Context) error {
	var body CreateBody
	err := c.Bind(&body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user, err := h.uc.CreateUser(c.Request().Context(), &usecase.CreateUserIn{Name: body.Name, Password: body.Password})
	if err != nil {
		return err
	}

	err = common.SetUserIDInCookie(c, user.ID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, CreateRes{})
	return nil
}
