package main

import (
	"github.com/labstack/echo/v4"
	"juliet/initializers"
	usertransport "juliet/internal/module/user/transport"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"message": "pong",
		})
	})

	v1 := e.Group("/v1")
	user := v1.Group("/user")
	{
		user.POST("/sign-up", usertransport.HandleSignUp(initializers.DB))
		user.GET("", usertransport.HandleFindUserByEmail(initializers.DB))
	}

	e.Logger.Fatal(e.Start(":1411"))
}
