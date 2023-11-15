package userstorage

import (
	"context"
	usermodel "juliet/internal/module/user/model"
)

func (s *postgresStorage) CreateNewUser(ctx context.Context, data *usermodel.SignUpModel) error {
	err := s.gorm.WithContext(ctx).Create(data).Error
	if err != nil {
		return err
	}

	return nil
}
