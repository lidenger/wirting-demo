package store

import (
	"gorm.io/gorm"
	"jwtenhance/model"
)

type SecretStore struct {
	DB *gorm.DB
}

func (s *SecretStore) SelectAllSecret() (result []*model.Secret, err error) {
	err = s.DB.Find(&result).Error
	return
}
