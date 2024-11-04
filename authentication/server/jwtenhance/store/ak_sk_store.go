package store

import (
	"gorm.io/gorm"
	"jwtenhance/model"
)

type AkSkStore struct {
	DB *gorm.DB
}

func (s *AkSkStore) SelectByAk(ak string) (result *model.AkSk, err error) {
	err = s.DB.First(&result, "ak = ?", ak).Error
	return
}

func (s *AkSkStore) SelectAll() (result []*model.AkSk, err error) {
	err = s.DB.Find(&result).Error
	return
}
