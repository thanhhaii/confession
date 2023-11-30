package userbusiness

import (
	"context"
	usermodel "juliet/internal/module/user/model"
)

type ListUserStore interface {
	ListUserByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) ([]usermodel.UserModel, error)
}

type listUserBiz struct {
	store ListUserStore
}

func NewListUserBusiness(store ListUserStore) *listUserBiz {
	return &listUserBiz{store: store}
}

func (biz *listUserBiz) ListUser(ctx context.Context) ([]usermodel.UserModel, error) {
	result, err := biz.store.ListUserByCondition(ctx, nil)
	if err != nil {
		return nil, err
	}

	return result, nil
}
