package model

import "time"

type Result[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type Order struct {
	ID          int64     `json:"id"`
	OrderNo     string    `json:"orderNo"`
	OrderStatus uint8     `json:"orderStatus"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
}
