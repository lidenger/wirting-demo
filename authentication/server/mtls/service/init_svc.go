package service

import (
	"gorm.io/gorm"
	"mtls/store"
)

var ServerSvcIns *ServerSvc
var OrderSvcIns *OrderSvc

func Initialize(db *gorm.DB) {
	serverStore := &store.ServerStore{DB: db}
	ServerSvcIns = &ServerSvc{store: serverStore}

	orderStore := &store.OrderStore{DB: db}
	OrderSvcIns = &OrderSvc{store: orderStore}

}
