package util

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"jwtbase/model"
	"testing"
	"time"
)

func TestJwtMarshal(t *testing.T) {
	// header
	header := &model.JwtHeader{}
	header.Algorithm = "HS256"
	header.Issuer = "Server2"
	// payload
	payload := &model.JwtPayload{}
	payload.Nonce = Generate16Str()
	payload.ServerSign = "s1"
	payload.Timestamp = time.Now().Unix()
	payload.SecretID = 1
	// signature
	key := []byte("4d8a8dcd8e8011ef91d48c32231f5813")
	data, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	sign, err := Signature(data, key, header.Algorithm)
	if err != nil {
		panic(err)
	}
	jwt := &model.Jwt{
		Header:    header,
		Payload:   payload,
		Signature: sign,
	}
	jwtStr, err := JwtMarshal(jwt)
	if err != nil {
		panic(err)
	}
	fmt.Println(jwtStr)
	fmt.Println("----unmarshal-----")
	jwtObj, err := JwtUnmarshal(jwtStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", jwtObj)
	fmt.Println("----verify sign-----")
	payload2, err := json.Marshal(jwt.Payload)
	if err != nil {
		panic(err)
	}
	sign2, err := Signature(payload2, key, jwt.Header.Algorithm)
	if err != nil {
		panic(err)
	}
	isVerify := hex.EncodeToString(sign) == hex.EncodeToString(sign2)
	fmt.Println("isVerify:", isVerify)
}
