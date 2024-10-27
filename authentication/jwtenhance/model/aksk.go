package model

import "time"

type AkSk struct {
	ID         int64     `gorm:"column:id;primary_key" json:"id"`
	ServerSign string    `gorm:"column:server_sign" json:"serverSign"`
	Ak         string    `gorm:"column:ak" json:"ak"`
	Sk         string    `gorm:"column:sk" json:"sk"`
	IsEnable   uint8     `gorm:"column:is_enable" json:"isEnable"`
	Remark     string    `gorm:"column:remark" json:"remark"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (s *AkSk) TableName() string {
	return "aksk"
}
