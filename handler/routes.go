package handler

import (
	"github.com/hayato24s/todo-echo-gorm/handler/auth"
	"github.com/hayato24s/todo-echo-gorm/handler/middleware"
	"github.com/hayato24s/todo-echo-gorm/handler/tasks"
	"github.com/hayato24s/todo-echo-gorm/handler/users"
	"github.com/hayato24s/todo-echo-gorm/usecase"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/hayato24s/todo-echo-gorm/docs"

	"github.com/labstack/echo/v4"
)

//	@title			Todo App Api
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func RegisterRoutes(e *echo.Echo, uc *usecase.UseCase) {

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(
		middleware.Auth,
		middleware.ErrorHandler,
	)

	auth.NewHandler(uc).Register(e)
	tasks.NewHandler(uc).Register(e)
	users.NewHandler(uc).Register(e)
}
