package service

import (
	"accesstoken/model"
	"accesstoken/store"
	"accesstoken/util"
	"errors"
	"fmt"
	"time"
)

// TOKEN_VALID_HOUR token有效期
const TOKEN_VALID_HOUR = 3 * 24 * time.Hour

type AccessTokenSvc struct {
	store   *store.AccessTokenStore
	akStore *store.AkSkStore
}

// 验证ak sk
func (s *AccessTokenSvc) verifyAkSk(serverSign, ak, sk string) (bool, error) {
	if ak == "" || sk == "" {
		return false, nil
	}
	aksk, err := s.akStore.SelectByAk(ak)
	if err != nil {
		return false, err
	}
	if aksk == nil {
		return false, nil
	}
	if aksk.Sk == sk && aksk.Sign == serverSign {
		return true, nil
	}
	return false, nil
}

// GenAccessToken 验证AK,SK生成AccessToken
func (s *AccessTokenSvc) GenAccessToken(serverSign, ak, sk string) (string, error) {
	verify, err := s.verifyAkSk(serverSign, ak, sk)
	if err != nil {
		return "", err
	}
	if !verify {
		return "", errors.New("无效的ak sk")
	}
	tokenM := &model.AccessToken{}
	tokenM.Sign = serverSign
	tokenM.Token = util.Generate32Str()
	now := time.Now()
	tokenM.CreateTime = now
	tokenM.UpdateTime = now
	tokenM.IsValid = 1
	err = s.store.Insert(tokenM)
	if err != nil {
		return "", err
	}
	return tokenM.Token, nil
}

// VerifyAccessToken 验证AccessToken
func (s *AccessTokenSvc) VerifyAccessToken(token string) (string, error) {
	if token == "" {
		return "", nil
	}
	tokenM, err := s.store.SelectByToken(token)
	if err != nil {
		return "", err
	}
	if tokenM == nil {
		return "", nil
	}
	// 已经无效,直接返回无效
	if tokenM.IsValid == 0 {
		return "", nil
	}
	// 有效,检测是否已失效
	if tokenM.IsValid == 1 {
		if tokenM.CreateTime.Add(TOKEN_VALID_HOUR).After(time.Now()) {
			return tokenM.Sign, nil
		}
		// 已无效，更新状态
		err = s.store.UpdateToExpire(token)
		if err != nil {
			return "", err
		}
		return "", nil
	}
	return "", errors.New(fmt.Sprintf("无效的Token状态: %d", tokenM.IsValid))
}
