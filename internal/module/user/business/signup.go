package userbusiness

import (
	"context"
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
	if err := biz.store.CreateNewUser(ctx, data); err != nil {
		return err
	}

	return nil
}
