package service

import (
	"accesstoken/store"
	"gorm.io/gorm"
)

var ServerSvcIns *ServerSvc
var AccessTokenSvcIns *AccessTokenSvc
var OrderSvcIns *OrderSvc

func Initialize(db *gorm.DB) {
	serverStore := &store.ServerStore{DB: db}
	ServerSvcIns = &ServerSvc{store: serverStore}

	accessTokenStore := &store.AccessTokenStore{DB: db}
	akSkStore := &store.AkSkStore{DB: db}

	AccessTokenSvcIns = &AccessTokenSvc{
		store:   accessTokenStore,
		akStore: akSkStore,
	}

	orderStore := &store.OrderStore{DB: db}
	OrderSvcIns = &OrderSvc{store: orderStore}

}
