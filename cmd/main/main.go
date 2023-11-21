package main

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"juliet/initializers"
	usertransport "juliet/internal/module/user/transport"
	"juliet/pkg/tokenFactory"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

type Service struct {
	tokenFactory tokenFactory.TokenFactory
}

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"message": "pong",
		})
	})

	service := Service{
		tokenFactory: tokenFactory.CreateTokenFactory(
			[]byte("secret-key"),
			jwt.SigningMethodHS256,
		),
	}

	v1 := e.Group("/v1")
	user := v1.Group("/user")
	{
		user.POST("/sign-up", usertransport.HandleSignUp(initializers.DB))
		user.GET("", usertransport.HandleFindUserByEmail(initializers.DB))
		user.POST("/sign-in", usertransport.HandleSignIn(initializers.DB, service.tokenFactory))
	}

	e.Logger.Fatal(e.Start(":1411"))
}
