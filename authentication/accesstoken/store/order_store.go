package store

import (
	"accesstoken/model"
	"gorm.io/gorm"
)

type OrderStore struct {
	DB *gorm.DB
}

func (s *OrderStore) SelectByNo(orderNo string) (result *model.Order, err error) {
	err = s.DB.First(&result, "order_no = ?", orderNo).Error
	return
}
