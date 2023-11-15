package initializers

import (
	usermodel "juliet/internal/module/user/model"
)

func SyncDatabase() {
	err := DB.AutoMigrate(&usermodel.UserModel{})
	if err != nil {
		panic("Failed when migrate table")
	}
}
