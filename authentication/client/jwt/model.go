package jwt_base

import "time"

type Order struct {
	ID          int64     `json:"id"`
	OrderNo     string    `json:"orderNo"`
	OrderStatus uint8     `json:"orderStatus"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
}

type AkSk struct {
	Ak string `json:"ak"`
	Sk string `json:"sk"`
}
