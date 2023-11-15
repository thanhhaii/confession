package userstorage

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"juliet/internal/module/user/model"
	"juliet/pkg/common"
)

func (s *postgresStorage) FindUser(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*usermodel.UserModel, error) {
	var (
		userRes usermodel.UserModel
		db      = s.gorm
	)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.WithContext(ctx).Where(conditions).First(&userRes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, err
	}

	return &userRes, nil
}
