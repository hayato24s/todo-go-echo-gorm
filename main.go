package main

import (
	"github.com/hayato24s/todo-echo-gorm/handler"
	"github.com/hayato24s/todo-echo-gorm/repository"
	"github.com/hayato24s/todo-echo-gorm/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	r, err := repository.NewRepository()
	if err != nil {
		panic(err)
	}
	uc := usecase.NewUseCase(r)
	handler.RegisterRoutes(e, uc)
	e.Logger.Fatal(e.Start(":1323"))
}
