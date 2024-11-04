package service

import (
	"gorm.io/gorm"
	"iplist/store"
)

var ServerSvcIns *ServerSvc

func Initialize(db *gorm.DB) {
	serverStore := &store.ServerStore{DB: db}
	ServerSvcIns = &ServerSvc{store: serverStore}
}
