package model

import "time"

type Server struct {
	ID         int64     `gorm:"column:id;primary_key" json:"id"`
	Sign       string    `gorm:"column:server_sign" json:"sign"`
	Name       string    `gorm:"column:server_name" json:"name"`
	Desc       string    `gorm:"column:server_desc" json:"desc"`
	IsEnable   uint8     `gorm:"column:is_enable" json:"isEnable"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

func (s *Server) TableName() string {
	return "server"
}
