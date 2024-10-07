package model

import "time"

type AccessToken struct {
	ID         int64     `gorm:"column:id;primary_key" json:"id"`
	Sign       string    `gorm:"column:server_sign" json:"sign"`
	Token      string    `gorm:"column:access_token" json:"token"`
	IsValid    uint8     `gorm:"column:is_valid" json:"isValid"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (s *AccessToken) TableName() string {
	return "access_token"
}
