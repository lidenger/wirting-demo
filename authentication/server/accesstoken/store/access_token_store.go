package store

import (
	"accesstoken/model"
	"gorm.io/gorm"
	"time"
)

type AccessTokenStore struct {
	DB *gorm.DB
}

// Insert 保存AccessToken信息
func (s *AccessTokenStore) Insert(token *model.AccessToken) (err error) {
	err = s.DB.Create(token).Error
	return
}

// UpdateToExpire 更新AccessToken标记为失效
func (s *AccessTokenStore) UpdateToExpire(token string) (err error) {
	return s.DB.Model(&model.AccessToken{}).
		Where("access_token = ?", token).
		Update("is_valid", 0).
		Update("update_time", time.Now()).Error
}

// SelectByToken 通过token获取AccessToken数据
func (s *AccessTokenStore) SelectByToken(token string) (result *model.AccessToken, err error) {
	err = s.DB.First(&result, "access_token = ?", token).Error
	return
}
