package userstorage

import "gorm.io/gorm"

type postgresStorage struct {
	gorm *gorm.DB
}

func NewPostgresStorage(db *gorm.DB) *postgresStorage {
	return &postgresStorage{gorm: db}
}
