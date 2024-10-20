package model

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"jwtbase/service"
	"jwtbase/util"
	"strings"
)

type JwtHeader struct {
	// 算法: HS256,HS512
	Algorithm string `json:"alg"`
	// 颁发者
	Issuer string `json:"iss"`
}

type JwtPayload struct {
	// 服务标识
	ServerSign string `json:"serverSign"`
	// 密钥ID
	SecretID int64 `json:"secretID"`
	// Nonce
	Nonce string `json:"nonce"`
	// 颁发时间
	Timestamp int64 `json:"timestamp"`
}

type Jwt struct {
	Header    *JwtHeader  `json:"header"`
	Payload   *JwtPayload `json:"payload"`
	Signature []byte      `json:"signature"`
}

func JwtMarshal(jwt *Jwt) (string, error) {
	headerPart, err := json.Marshal(jwt.Header)
	if err != nil {
		return "", err
	}
	payloadPart, err := json.Marshal(jwt.Payload)
	if err != nil {
		return "", err
	}
	result := base64.RawURLEncoding.EncodeToString(headerPart) + "." +
		base64.RawURLEncoding.EncodeToString(payloadPart) + "." +
		base64.RawURLEncoding.EncodeToString(jwt.Signature)
	return result, nil
}

func JwtUnmarshal(jwt string) (*Jwt, error) {
	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid jwt format")
	}
	headerPart, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, err
	}
	jwtHeader := &JwtHeader{}
	err = json.Unmarshal(headerPart, jwtHeader)
	if err != nil {
		return nil, err
	}
	payloadPart, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	jwtPayload := &JwtPayload{}
	err = json.Unmarshal(payloadPart, jwtPayload)
	if err != nil {
		return nil, err
	}
	secret := service.FetchSecretByID(jwtPayload.SecretID)
	// 验签
	sign, err := util.Signature(payloadPart, secret.Secret)
	if err != nil {
		return nil, err
	}
	if sign != parts[2] {
		return nil, errors.New("invalid jwt signature")
	}
	signBytes, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, err
	}
	return &Jwt{
		Header:    jwtHeader,
		Payload:   jwtPayload,
		Signature: signBytes,
	}, nil
}
