package service

import (
	"gorm.io/gorm"
	"jwtbase/store"
)

var ServerSvcIns *ServerSvc
var OrderSvcIns *OrderSvc
var JwtSvcIns *JwtSvc
var AkSkSvcIns *AkSkSvc

func Initialize(db *gorm.DB) {
	serverStore := &store.ServerStore{DB: db}
	ServerSvcIns = &ServerSvc{store: serverStore}

	orderStore := &store.OrderStore{DB: db}
	OrderSvcIns = &OrderSvc{store: orderStore}

	secretStore := &store.SecretStore{DB: db}
	JwtSvcIns = &JwtSvc{store: secretStore}

	akskStore := &store.AkSkStore{DB: db}
	AkSkSvcIns = &AkSkSvc{store: akskStore}

}
