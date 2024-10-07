package request_sign

import "testing"

func TestGetOrderByNo(t *testing.T) {
	order, err := GetOrderByNo("425c3b29832711ef8f0f8c32")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", order)
}

func TestGenSignatureHeader(t *testing.T) {
	url := "/v1/order/425c3b29832711ef8f0f8c32"
	auth := GenSignatureHeader(url)
	t.Log(auth)
}
