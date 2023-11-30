package usermodel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserProfile struct {
	FirstName string `json:"username"`
	LastName  string `json:"lastName"`
}

func (u *UserProfile) Value() (driver.Value, error) {
	byteData, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	return string(byteData), nil
}

func (u *UserProfile) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	err := json.Unmarshal(bytes, &u)
	return err
}

type UserModel struct {
	gorm.Model
	Email    string      `json:"email" gorm:"unique"`
	Password string      `json:"-"`
	Profile  UserProfile `json:"profile"`
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
