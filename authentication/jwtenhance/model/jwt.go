package model

type JwtHeader struct {
	// 算法: HS256,HS512
	Algorithm string `json:"alg"`
	// 颁发者
	Issuer string `json:"iss"`
}

type JwtPayload struct {
	// JWT ID
	JwtID string `json:"jwtID"`
	// AK
	AK string `json:"ak"`
	// 服务标识
	ServerSign string `json:"serverSign"`
	// 密钥ID
	SecretID int64 `json:"secretID"`
	// Nonce
	Nonce string `json:"nonce"`
	// IP
	IP string `json:"ip"`
	// 颁发时间
	Timestamp int64 `json:"timestamp"`
}

type Jwt struct {
	Header    *JwtHeader  `json:"header"`
	Payload   *JwtPayload `json:"payload"`
	Signature []byte      `json:"signature"`
}
