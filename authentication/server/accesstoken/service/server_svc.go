package service

import (
	"accesstoken/model"
	"accesstoken/store"
	"errors"
)

type ServerSvc struct {
	store *store.ServerStore
}

func (s *ServerSvc) GetBySign(sign string) (*model.Server, error) {
	if sign == "" {
		return nil, errors.New("sign不能为空")
	}
	return s.store.SelectServerBySign(sign)
}
