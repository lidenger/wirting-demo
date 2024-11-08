package service

import (
	"errors"
	"mtls/model"
	"mtls/store"
)

type ServerSvc struct {
	store *store.ServerStore
}

func (s *ServerSvc) GetBySign(sign string) (*model.Server, error) {
	if sign == "" {
		return nil, errors.New("sign不能为空")
	}
	return s.store.SelectBySign(sign)
}
