package service

import (
	"jwtenhance/model"
	"jwtenhance/store"
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

func (s *AkSkSvc) IsDisable(ak string) bool {
	sk := akCache[ak]
	if sk == nil {
		return false
	}
	return sk.IsEnable == 0

}
