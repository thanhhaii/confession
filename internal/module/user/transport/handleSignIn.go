package usertransport

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"juliet/pkg/common"
	"juliet/pkg/tokenFactory"
	"net/http"

	userbusiness "juliet/internal/module/user/business"
	usermodel "juliet/internal/module/user/model"
	userstorage "juliet/internal/module/user/storage"
)

func HandleSignIn(db *gorm.DB, tokenService tokenFactory.TokenFactory) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payloadSignUp usermodel.SignInModel
			err           error
			token         string
		)

		if err = c.Bind(&payloadSignUp); err != nil {
			return c.JSON(http.StatusBadRequest, common.Response[bool]{
				Success: false,
				Status:  400,
				Message: err.Error(),
			})
		}

		storage := userstorage.NewPostgresStorage(db)
		business := userbusiness.NewSignInStorage(storage, tokenService)

		token, err = business.HandleSignIn(c.Request().Context(), &payloadSignUp)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.Response[bool]{
				Success: false,
				Status:  500,
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, common.Response[string]{
			Success: true,
			Status:  201,
			Data:    token,
		})
	}
}
