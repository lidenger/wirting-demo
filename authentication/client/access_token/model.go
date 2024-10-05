package access_token

import "time"

type Order struct {
	ID          int64     `json:"id"`
	OrderNo     string    `json:"orderNo"`
	OrderStatus uint8     `json:"orderStatus"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
}
