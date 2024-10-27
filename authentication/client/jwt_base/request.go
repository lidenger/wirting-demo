package jwt_base

import (
	"bytes"
	"client/model"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	Domain     = "http://127.0.0.1"
	ServerSign = "s1"
	AK         = "b4a9bcc1825f11ef"
	SK         = "c26c6f9c825f11efa5ae8c32"
)

func GetOrderByNo(orderNo string) (*Order, error) {
	url := "/v1/order/" + orderNo
	req, err := http.NewRequest("GET", Domain+url, nil)
	if err != nil {
		return nil, err
	}
	jwt, err := GetJwt()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", jwt)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return analysisResp[*Order](resp)
}

func GetJwt() (string, error) {
	url := Domain + "/genToken"
	param := &AkSk{
		Ak: AK,
		Sk: SK,
	}
	jsonBytes, err := json.Marshal(param)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonBytes))
	if err != nil {
		return "", err
	}
	jwt, err := analysisResp[string](resp)
	return jwt, err
}

func analysisResp[T any](resp *http.Response) (T, error) {
	d := *new(T)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return d, err
	}
	if resp.StatusCode != 200 {
		result := &model.Result[any]{}
		err = json.Unmarshal(body, result)
		if err != nil {
			return d, err
		}
		return d, errors.New(result.Msg)
	}
	result := &model.Result[T]{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return d, err
	}
	if result.Code < 0 {
		return d, errors.New(result.Msg)
	}
	return result.Data, nil
}
