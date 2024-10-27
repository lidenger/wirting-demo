package jwt_base

import "testing"

func TestGetOrderByNo(t *testing.T) {
	order, err := GetOrderByNo("425c3b29832711ef8f0f8c32")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", order)
}

func TestGetJwt(t *testing.T) {
	jwt, err := GetJwt()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", jwt)
}
