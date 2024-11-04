package service

import (
	"mtls/model"
	"mtls/store"
)

type OrderSvc struct {
	store *store.OrderStore
}

func (s *OrderSvc) FetchOrderByNo(orderNo string) (*model.Order, error) {
	return s.store.SelectByNo(orderNo)
}
