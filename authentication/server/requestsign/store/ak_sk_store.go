package store

import (
	"gorm.io/gorm"
	"requestsign/model"
)

type AkSkStore struct {
	DB *gorm.DB
}

func (s *AkSkStore) SelectByAk(ak string) (result *model.AkSk, err error) {
	err = s.DB.First(&result, "ak = ?", ak).Error
	return
}
