package middleware

import (
	"requestsign/config"
	"requestsign/service"
	"requestsign/store"
	"testing"
)

func TestMain(m *testing.M) {
	conf := config.Initialize()
	db := store.Initialize(conf)
	service.Initialize(db)
	m.Run()
}

func TestAnalysisHeaderSignature(t *testing.T) {
	authorization := "algorithm=HMAC-SHA256,ak=b4a9bcc1825f11ef,time=1728222161,nonce=29b69dd683e811ef,signature=313233616263746573743132336a6a6b6a3b6c6b6a6b6c3b617364"
	params := AnalysisHeaderSignature(authorization)
	t.Logf("%+v", params)
}

func TestSignature(t *testing.T) {
	algorithm := CheckAlgorithm("HMAC-SHA256")
	ak := "b4a9bcc1825f11ef"
	sk := "c26c6f9c825f11efa5ae8c32"
	nonce := "29b69dd683e811ef"
	time := "1728222161"
	url := "/v1/order/425c3b29832711ef8f0f8c32"
	sign := Signature(algorithm, ak, sk, nonce, time, url)
	t.Log(sign)
}

func TestVerifySignature(t *testing.T) {
	authorization := "algorithm=HMAC-SHA256,ak=b4a9bcc1825f11ef,time=1728222161,nonce=29b69dd683e811ef,signature=f7171365d167239e93d004d87763f0618a82ae141a7c1a65e34b53ef2b3c765b"
	url := "/v1/order/425c3b29832711ef8f0f8c32"
	sk := "c26c6f9c825f11efa5ae8c32"
	params := AnalysisHeaderSignature(authorization)
	err := VerifySignature(params, sk, url)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("验证通过")
	}
}
