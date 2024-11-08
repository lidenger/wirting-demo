package model

import "time"

type JwtDeactive struct {
	ID         int64     `gorm:"column:id;primary_key" json:"id"`
	JwtID      string    `gorm:"column:jwt_id" json:"jwtID"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
}

func (s *JwtDeactive) TableName() string {
	return "jwt_deactive"
}
