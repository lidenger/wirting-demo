package store

import (
	"gorm.io/gorm"
	"jwtenhance/model"
)

type JwtDeactiveStore struct {
	DB *gorm.DB
}

func (s *JwtDeactiveStore) SelectAll() (result []*model.JwtDeactive, err error) {
	err = s.DB.Find(&result).Error
	return
}
