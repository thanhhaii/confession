package usertransport

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	userbusiness "juliet/internal/module/user/business"
	usermodel "juliet/internal/module/user/model"
	userstorage "juliet/internal/module/user/storage"
	"juliet/pkg/common"
	"net/http"
)

func HandleGetListUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		storage := userstorage.NewPostgresStorage(db)
		business := userbusiness.NewListUserBusiness(storage)

		result, err := business.ListUser(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusOK, common.Response[string]{
				Success: false,
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, common.Response[[]usermodel.UserModel]{
			Success: true,
			Status:  200,
			Data:    result,
		})
	}
}
