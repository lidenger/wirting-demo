package service

import (
	"errors"
	"jwtenhance/model"
	"jwtenhance/store"
)

var serverCache = make(map[string]*model.Server)

type ServerSvc struct {
	store *store.ServerStore
}

func (s *ServerSvc) GetBySign(sign string) (*model.Server, error) {
	server := serverCache[sign]
	if server == nil {
		return nil, errors.New("server not exist")
	}
	return server, nil
}

func (s *ServerSvc) LoadAll() {
	servers, err := s.store.SelectAll()
	if err != nil {
		panic(err)
	}
	for _, server := range servers {
		serverCache[server.Sign] = server
	}
}

func (s *ServerSvc) IsDisable(sign string) bool {
	server := serverCache[sign]
	if server == nil {
		return false
	}
	return server.IsEnable == 0
}
