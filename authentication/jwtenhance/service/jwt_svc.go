package service

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"google.golang.org/protobuf/proto"
	"jwtenhance/model"
	"jwtenhance/protogen"
	"jwtenhance/store"
	"jwtenhance/util"
	"math/rand/v2"
	"time"
)

type JwtSvc struct {
	store         *store.SecretStore
	deactiveStore *store.JwtDeactiveStore
}

var allSecretCache = make(map[int64]*model.Secret)
var deactiveCache = make(map[string]struct{})
var allID = make([]int64, 0)

func FetchSecretByID(id int64) *model.Secret {
	return allSecretCache[id]
}

func (s *JwtSvc) LoadAllSecret() {
	arr, err := s.store.SelectAllSecret()
	if err != nil {
		panic(err)
	}
	for _, secret := range arr {
		if secret.IsEnable == 1 {
			allSecretCache[secret.ID] = secret
			allID = append(allID, secret.ID)
		}
	}
	if len(allSecretCache) == 0 {
		panic("没有可用的系统密钥")
	}
}

func RandomSecret() *model.Secret {
	n := rand.IntN(len(allID))
	ID := allID[n]
	return allSecretCache[ID]
}

func GenJwt(serverSign, ip string) (string, error) {
	// header
	header := &protogen.JwtHeader{}
	header.Algorithm = "HS256"
	header.Issuer = "Server2"
	// payload
	payload := &protogen.JwtPayload{}
	payload.JwtId = util.Generate16Str()
	payload.Ip = ip
	payload.Nonce = util.Generate16Str()
	payload.ServerSign = serverSign
	payload.Timestamp = time.Now().Unix()
	secret := RandomSecret()
	payload.SecretId = secret.ID
	// signature
	key := []byte(secret.Secret)
	data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	sign, err := util.Signature(data, key, header.Algorithm)
	if err != nil {
		return "", err
	}
	jwt := &protogen.Jwt{
		Header:    header,
		Payload:   payload,
		Signature: sign,
	}
	jwtData, err := proto.Marshal(jwt)
	if err != nil {
		return "", err
	}
	// Encrypt AES/CBC/PKCS#7
	secret1 := FetchSecretByID(1)
	cipher, err := util.Encrypt([]byte(secret1.Secret), jwtData)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(cipher), nil
}

func VerifyJwt(jwtStr string) (string, error) {
	// Decrypt AES/CBC/PKCS#7
	secret1 := FetchSecretByID(1)
	jwtDecode, err := hex.DecodeString(jwtStr)
	if err != nil {
		return "", err
	}
	jwtData, err := util.Decrypt([]byte(secret1.Secret), jwtDecode)
	jwt := &protogen.Jwt{}
	err = proto.Unmarshal(jwtData, jwt)
	if err != nil {
		return "", err
	}
	// 验证单个JWT是否被注销
	if _, exists := deactiveCache[jwt.Payload.JwtId]; exists {
		return "", errors.New("非法JWT")
	}
	// 验证AK是否禁用
	if AkSkSvcIns.IsDisable(jwt.Payload.Ak) {
		return "", errors.New("非法JWT")
	}
	// 验证服务是否禁用
	if ServerSvcIns.IsDisable(jwt.Payload.ServerSign) {
		return "", errors.New("非法JWT")
	}
	payload, err := json.Marshal(jwt.Payload)
	if err != nil {
		return "", err
	}
	secret := FetchSecretByID(jwt.Payload.SecretId)
	key := []byte(secret.Secret)
	header := jwt.Header
	sign, err := util.Signature(payload, key, header.Algorithm)
	if err != nil {
		return "", err
	}
	isVerify := hex.EncodeToString(sign) == hex.EncodeToString(jwt.Signature)
	if isVerify {
		return jwt.Payload.ServerSign, nil
	} else {
		return "", errors.New("验签失败，非法JWT")
	}
}

// LoadAllDeactiveJwt 加载所有注销的JWT
func (s *JwtSvc) LoadAllDeactiveJwt() {
	data, err := s.deactiveStore.SelectAll()
	if err != nil {
		panic(err)
	}
	for _, d := range data {
		deactiveCache[d.JwtID] = struct{}{}
	}
}
