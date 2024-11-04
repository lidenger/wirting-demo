package mtls

import (
	"client/model"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

func GetOrderByNo(orderNo string) (*model.Order, error) {
	domain := "https://local.com"
	url := "/v1/order/" + orderNo

	// 配置server ca证书
	pool := x509.NewCertPool()
	caCrt, err := os.ReadFile("serverroot.crt")
	if err != nil {
		return nil, err
	}
	pool.AppendCertsFromPEM(caCrt)
	s1Crt, err := tls.LoadX509KeyPair("s1.crt", "s1.key")
	if err != nil {
		return nil, err
	}
	// 配置s1证书
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{s1Crt},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(domain + url)
	if err != nil {
		return nil, err
	}
	return analysisResp[*model.Order](resp)
}

func analysisResp[T any](resp *http.Response) (T, error) {
	d := *new(T)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return d, err
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
