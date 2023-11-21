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
	return UserModel{}.TableName()
}

type SignInModel struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (SignInModel) TableName() string {
	return UserModel{}.TableName()
}
