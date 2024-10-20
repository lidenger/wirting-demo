package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"jwtbase/model"
	"jwtbase/store"
	"jwtbase/util"
	"math/rand/v2"
	"time"
)

type JwtSvc struct {
	store *store.SecretStore
}

var allSecretMap = make(map[int64]*model.Secret)
var allID = make([]int64, 0)

func FetchSecretByID(id int64) *model.Secret {
	return allSecretMap[id]
}

func (s *JwtSvc) LoadAllSecret() error {
	arr, err := s.store.SelectAllSecret()
	if err != nil {
		return err
	}
	for _, secret := range arr {
		if secret.IsEnable == 1 {
			allSecretMap[secret.ID] = secret
			allID = append(allID, secret.ID)
		}
	}
	if len(allSecretMap) == 0 {
		panic("没有可用的系统密钥")
	}
	return nil
}

func (s *JwtSvc) RandomSecret() *model.Secret {
	n := rand.IntN(len(allID))
	ID := allID[n]
	return allSecretMap[ID]
}

func (s *JwtSvc) GenJwt(serverSign string) (string, error) {
	// header
	header := &model.JwtHeader{}
	header.Algorithm = "HS256"
	header.Issuer = "Server2"
	// payload
	payload := &model.JwtPayload{}
	payload.Nonce = util.Generate16Str()
	payload.ServerSign = serverSign
	payload.Timestamp = time.Now().Unix()
	secret := s.RandomSecret()
	payload.SecretID = secret.ID
	// signature
	key := []byte(secret.Secret)
	h := hmac.New(sha256.New, key)
	data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	h.Write(data)
	//digested := h.Sum(nil)
	//
}
