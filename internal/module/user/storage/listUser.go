package userstorage

import (
	"context"
	usermodel "juliet/internal/module/user/model"
)

func (s *postgresStorage) ListUserByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) ([]usermodel.UserModel, error) {
	var (
		result []usermodel.UserModel
		db     = s.gorm
	)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.WithContext(ctx).Table(usermodel.UserModel{}.TableName()).
		Where(conditions).Find(&result)

	return result, nil
}
