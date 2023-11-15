package userbusiness

import (
	"context"
	usermodel "juliet/internal/module/user/model"
)

type FindUserStorage interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*usermodel.UserModel, error)
}

type findUserBiz struct {
	store FindUserStorage
}

func NewFindUserStorage(store FindUserStorage) *findUserBiz {
	return &findUserBiz{store: store}
}

func (biz *findUserBiz) FindUserByEmail(ctx context.Context, email string) (*usermodel.UserModel, error) {
	return biz.store.FindUser(ctx, map[string]interface{}{
		"email": email,
	})
}
