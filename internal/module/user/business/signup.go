package userbusiness

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	usermodel "juliet/internal/module/user/model"
)

type SignUpStorage interface {
	CreateNewUser(ctx context.Context, data *usermodel.SignUpModel) error
}

type signUpBiz struct {
	store SignUpStorage
}

func NewSignUpStorage(store SignUpStorage) *signUpBiz {
	return &signUpBiz{store: store}
}

func (biz *signUpBiz) SignUpNewAccount(ctx context.Context, data *usermodel.SignUpModel) error {
	var err error
	var passwordHash []byte

	passwordHash, err = bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		return err
	}

	data.Password = string(passwordHash)
	if err := biz.store.CreateNewUser(ctx, data); err != nil {
		return err
	}

	return nil
}
