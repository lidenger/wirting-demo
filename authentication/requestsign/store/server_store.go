package store

import (
	"gorm.io/gorm"
	"requestsign/model"
)

type ServerStore struct {
	DB *gorm.DB
}

// SelectServerBySign 通过服务标识获取服务信息
func (s *ServerStore) SelectServerBySign(sign string) (result *model.Server, err error) {
	err = s.DB.First(&result, "server_sign = ?", sign).Error
	return
}
