package router

import (
	"net/http"

	"sample/controller"

	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/user/:id", controller.GetUser)
	e.POST("/user", controller.CreateUser)
	e.PUT("/user/:id", controller.UpdateUser)
	e.DELETE("/user/:id", controller.DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}
