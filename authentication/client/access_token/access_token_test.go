package access_token

import "testing"

func TestGetAccessToken(t *testing.T) {
	accessToken, err := GetAccessToken()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(accessToken)
}

func TestGetOrderByNo(t *testing.T) {
	order, err := GetOrderByNo("425c3b29832711ef8f0f8c32")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(order)
}
