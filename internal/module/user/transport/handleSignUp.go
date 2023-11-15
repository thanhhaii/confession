package usertransport

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"juliet/pkg/common"
	"net/http"

	userbusiness "juliet/internal/module/user/business"
	usermodel "juliet/internal/module/user/model"
	userstorage "juliet/internal/module/user/storage"
)

func HandleSignUp(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payloadSignUp usermodel.SignUpModel

		if err := c.Bind(&payloadSignUp); err != nil {
			return c.JSON(http.StatusBadRequest, common.Response[bool]{
				Success: false,
				Status:  400,
				Message: err.Error(),
			})
		}

		storage := userstorage.NewPostgresStorage(db)
		business := userbusiness.NewSignUpStorage(storage)

		if err := business.SignUpNewAccount(c.Request().Context(), &payloadSignUp); err != nil {
			return c.JSON(http.StatusInternalServerError, common.Response[bool]{
				Success: false,
				Status:  500,
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, common.Response[bool]{
			Success: true,
			Status:  201,
		})
	}
}
