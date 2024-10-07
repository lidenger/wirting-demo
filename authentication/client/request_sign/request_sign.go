package request_sign

import (
	"client/model"
	"client/util"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetOrderByNo(orderNo string) (Order, error) {
	domain := "http://127.0.0.1"
	url := "/v1/order/" + orderNo
	req, err := http.NewRequest("GET", domain+url, nil)
	if err != nil {
		return Order{}, err
	}
	req.Header.Set("Authorization", GenSignatureHeader(url))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Order{}, err
	}
	return analysisResp[Order](resp)
}

// GenSignatureHeader 生成Authorization
func GenSignatureHeader(url string) string {
	algorithm := sha256.New
	ak := "b4a9bcc1825f11ef"
	sk := "c26c6f9c825f11efa5ae8c32"
	nonce := util.Generate16Str()
	t := time.Now().Unix()
	ts := strconv.Itoa(int(t))
	sign := Signature(algorithm, ak, sk, nonce, ts, url)
	return fmt.Sprintf("algorithm=%s,ak=%s,time=%s,nonce=%s,signature=%s", "HMAC-SHA256", ak, ts, nonce, sign)
}

// Signature 签名
func Signature(hash func() hash.Hash, ak, sk, nonce, time, url string) string {
	dataBuilder := &strings.Builder{}
	dataBuilder.WriteString(ak)
	dataBuilder.WriteString(nonce)
	dataBuilder.WriteString(time)
	dataBuilder.WriteString(url)
	data := dataBuilder.String()
	h := hmac.New(hash, []byte(sk))
	h.Write([]byte(data))
	digested := h.Sum(nil)
	return hex.EncodeToString(digested)
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
