package store

import (
	"gorm.io/gorm"
	"mtls/model"
)

type ServerStore struct {
	DB *gorm.DB
}

func (s *ServerStore) SelectByIp(ip string) (result *model.ServerList, err error) {
	err = s.DB.First(&result, "ip = ?", ip).Error
	return
}

func (s *ServerStore) SelectBySign(sign string) (result *model.Server, err error) {
	err = s.DB.First(&result, "server_sign = ?", sign).Error
	return
}
