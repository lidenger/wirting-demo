package main

import (
	"bytes"
	"client/access_token"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	err := accessToken()
	if err != nil {
		panic(err)
	}
}

// server1使用accessToken方式完成认证
func accessToken() error {
	// 获取accessToken
	url := "http://127.0.0.1/access_token"
	param := &access_token.GenAccessTokenParam{
		Sign: "s1",
		Ak:   "b4a9bcc1825f11ef",
		Sk:   "c26c6f9c825f11efa5ae8c32",
	}
	json, err := json.Marshal(param)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(json))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
