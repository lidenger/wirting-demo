package service

import (
	"gorm.io/gorm"
	"requestsign/store"
)

var ServerSvcIns *ServerSvc
var AkSkSvcIns *AkSkSvc
var OrderSvcIns *OrderSvc

func Initialize(db *gorm.DB) {
	serverStore := &store.ServerStore{DB: db}
	ServerSvcIns = &ServerSvc{store: serverStore}

	akSkStore := &store.AkSkStore{DB: db}
	AkSkSvcIns = &AkSkSvc{store: akSkStore}

	orderStore := &store.OrderStore{DB: db}
	OrderSvcIns = &OrderSvc{store: orderStore}

}
