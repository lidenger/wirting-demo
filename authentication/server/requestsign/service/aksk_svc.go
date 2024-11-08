package service

import (
	"requestsign/model"
	"requestsign/store"
)

type AkSkSvc struct {
	store *store.AkSkStore
}

func (s *AkSkSvc) GetByAk(ak string) (*model.AkSk, error) {
	return s.store.SelectByAk(ak)
}
