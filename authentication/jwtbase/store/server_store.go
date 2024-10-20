package store

import (
	"gorm.io/gorm"
	"jwtbase/model"
)

type ServerStore struct {
	DB *gorm.DB
}

func (s *ServerStore) SelectBySign(sign string) (result *model.Server, err error) {
	err = s.DB.First(&result, "server_sign = ?", sign).Error
	return
}
