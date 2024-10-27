package util

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"jwtenhance/model"
	"strings"
)

func JwtMarshal(jwt *model.Jwt) (string, error) {
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

func JwtUnmarshal(jwt string) (*model.Jwt, error) {
	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid jwt format")
	}
	headerPart, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, err
	}
	jwtHeader := &model.JwtHeader{}
	err = json.Unmarshal(headerPart, jwtHeader)
	if err != nil {
		return nil, err
	}
	payloadPart, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	jwtPayload := &model.JwtPayload{}
	err = json.Unmarshal(payloadPart, jwtPayload)
	if err != nil {
		return nil, err
	}
	signBytes, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, err
	}
	return &model.Jwt{
		Header:    jwtHeader,
		Payload:   jwtPayload,
		Signature: signBytes,
	}, nil
}
