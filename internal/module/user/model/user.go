package usermodel

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (UserModel) TableName() string {
	return "users"
}

type SignUpModel struct {
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (SignUpModel) TableName() string {
	return "users"
}
