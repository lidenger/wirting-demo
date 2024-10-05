package access_token

import (
	"bytes"
	"client/model"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetAccessToken() (string, error) {
	url := "http://127.0.0.1/access_token"
	param := &GenAccessTokenParam{
		Sign: "s1",
		Ak:   "b4a9bcc1825f11ef",
		Sk:   "c26c6f9c825f11efa5ae8c32",
	}
	jsonBytes, err := json.Marshal(param)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonBytes))
	if err != nil {
		return "", err
	}
	return analysisResp[string](resp)
}

func GetOrderByNo(orderNo string) (Order, error) {
	token, err := GetAccessToken()
	if err != nil {
		return Order{}, err
	}
	url := "http://127.0.0.1/v1/order/" + orderNo
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Order{}, err
	}
	req.Header.Set("accessToken", token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Order{}, err
	}
	return analysisResp[Order](resp)
}

func analysisResp[T any](resp *http.Response) (T, error) {
	d := *new(T)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return d, err
	}
	log.Println(string(body))
	result := &model.Result[T]{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return d, err
	}
	return result.Data, nil
}
