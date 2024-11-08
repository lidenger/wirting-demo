package model

import "time"

type ServerList struct {
	ID         int64     `gorm:"column:id;primary_key" json:"id"`
	Sign       string    `gorm:"column:server_sign" json:"serverSign"`
	Ip         string    `gorm:"column:ip" json:"ip"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

func (s *ServerList) TableName() string {
	return "server_ip"
}
