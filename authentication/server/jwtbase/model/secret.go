package model

import "time"

type Secret struct {
	ID         int64     `gorm:"column:id;primary_key" json:"id"`
	Secret     string    `gorm:"column:secret" json:"secret"`
	IsEnable   uint8     `gorm:"column:is_enable" json:"isEnable"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

func (s *Secret) TableName() string {
	return "sys_secret"
}
