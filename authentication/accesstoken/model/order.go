package model

import "time"

type Order struct {
	ID          int64     `gorm:"column:id;primary_key" json:"id"`
	OrderNo     string    `gorm:"column:order_no" json:"orderNo"`
	OrderStatus uint8     `gorm:"column:order_status" json:"orderStatus"`
	CreateTime  time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (s *Order) TableName() string {
	return "order_info"
}
