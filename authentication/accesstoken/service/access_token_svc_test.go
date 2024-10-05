package service

import (
	"accesstoken/config"
	"accesstoken/store"
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	conf := config.Initialize()
	db := store.Initialize(conf)
	Initialize(db)
	m.Run()
}

func TestGetAccessToken(t *testing.T) {
	serverSign := "s1"
	ak := "b4a9bcc1825f11ef"
	sk := "c26c6f9c825f11efa5ae8c32"
	token, err := AccessTokenSvcIns.GenAccessToken(serverSign, ak, sk)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)
}

func TestVerifyAccessToken(t *testing.T) {
	token := "8d9ca373830011efbf8b8c32231f5813"
	sign, err := AccessTokenSvcIns.VerifyAccessToken(token)
	if err != nil {
		t.Fatal(err)
	}
	if sign == "" {
		t.Fatal("无效token")
	}
	fmt.Println(sign)
}
