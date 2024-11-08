package store

import (
	"gorm.io/gorm"
	"jwtenhance/model"
)

type ServerStore struct {
	DB *gorm.DB
}

func (s *ServerStore) SelectBySign(sign string) (result *model.Server, err error) {
	err = s.DB.First(&result, "server_sign = ?", sign).Error
	return
}

func (s *ServerStore) SelectAll() (result []*model.Server, err error) {
	err = s.DB.Find(&result).Error
	return
}
