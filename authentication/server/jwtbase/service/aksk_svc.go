package service

import (
	"jwtbase/model"
	"jwtbase/store"
)

var akCache map[string]*model.AkSk

type AkSkSvc struct {
	store *store.AkSkStore
}

func (s *AkSkSvc) GetSk(ak string) *model.AkSk {
	return akCache[ak]
}

func (s *AkSkSvc) LoadAllAkSk() {
	akCache = make(map[string]*model.AkSk)
	ms, err := s.store.SelectAll()
	if err != nil {
		panic(err)
	}
	for _, m := range ms {
		akCache[m.Ak] = m
	}
}
