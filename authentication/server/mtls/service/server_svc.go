package service

import (
	"mtls/model"
	"mtls/store"
)

type ServerSvc struct {
	store *store.ServerStore
}

func (s *ServerSvc) GetServerByIp(ip string) (*model.Server, error) {
	serverIp, err := s.store.SelectByIp(ip)
	if err != nil {
		return nil, err
	}
	if serverIp == nil {
		return nil, nil
	}
	return s.store.SelectBySign(serverIp.Sign)
}
