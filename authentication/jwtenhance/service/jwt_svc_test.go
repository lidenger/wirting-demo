package service

import (
	"jwtenhance/model"
	"testing"
)

func TestMain(m *testing.M) {
	allSecretCache[1] = &model.Secret{
		ID:     1,
		Secret: "4d8a8dcd8e8011ef91d48c32231f5813",
	}
	allID = append(allID, allSecretCache[1].ID)
	m.Run()
}

func TestJwtSvc_GenJwt(t *testing.T) {
	jwt, err := GenJwt("s1xxx", "10.0.0.1")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", jwt)
	t.Logf("len:%d", len(jwt))
}
