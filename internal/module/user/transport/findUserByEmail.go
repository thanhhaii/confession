package usertransport

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"

	userbusiness "juliet/internal/module/user/business"
	usermodel "juliet/internal/module/user/model"
	userstorage "juliet/internal/module/user/storage"
)

func HandleFindUserByEmail(db *gorm.DB) echo.HandlerFunc {
	type Query struct {
		Email string `query:"email"`
	}

	return func(c echo.Context) error {
		var (
			query Query
			err   error
			user  *usermodel.UserModel
		)
		if err := c.Bind(&query); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
		}

		storage := userstorage.NewPostgresStorage(db)
		business := userbusiness.NewFindUserStorage(storage)

		if user, err = business.FindUserByEmail(c.Request().Context(), query.Email); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"data":    user,
		})
	}
}
