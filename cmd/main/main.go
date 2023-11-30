package main

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"juliet/initializers"
	usertransport "juliet/internal/module/user/transport"
	"juliet/pkg/tokenFactory"
	"strings"
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

	server := e.Group("/server")
	{
		server.GET("/ping", func(c echo.Context) error {
			return c.JSON(200, map[string]interface{}{
				"message": "pong",
			})
		})
	}

	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte("secret-key"),
		SigningMethod: echojwt.AlgorithmHS256,
		Skipper: func(c echo.Context) bool {
			path := c.Request().URL.Path
			return strings.Contains(path, "/sign-in") || strings.Contains(path, "/sign-up")
		},
	}))

	service := Service{
		tokenFactory: tokenFactory.CreateTokenFactory(
			[]byte("secret-key"),
			jwt.SigningMethodHS256,
		),
	}

	v1 := e.Group("/v1")
	user := v1.Group("/user")
	{
		user.GET("", usertransport.HandleFindUserByEmail(initializers.DB))
		user.GET("/list", usertransport.HandleGetListUser(initializers.DB))
		user.POST("/sign-up", usertransport.HandleSignUp(initializers.DB))
		user.POST("/sign-in", usertransport.HandleSignIn(initializers.DB, service.tokenFactory))
	}

	e.Logger.Fatal(e.Start(":1411"))
}
