package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"hash"
	"log"
	"requestsign/service"
	"strconv"
	"strings"
	"time"
)

// VALID_SECOND 签名有效期
const VALID_SECOND = 20 * 60 * time.Second

func Authentication(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		c.Next()
		return
	}
	params := AnalysisHeaderSignature(authorization)
	aksk, err := service.AkSkSvcIns.GetByAk(params["ak"])
	if err != nil {
		log.Printf("+%v", err)
		c.Next()
		return
	}
	url := c.Request.URL.String()
	// 验证有效期
	t, err := strconv.Atoi(params["time"])
	if err != nil {
		log.Printf("+%v", err)
		c.Next()
		return
	}
	if time.Unix(int64(t), 0).Add(VALID_SECOND).Before(time.Now()) {
		log.Printf("authorization已过期：" + time.Unix(int64(t), 0).Format(time.DateTime))
		c.Next()
		return
	}
	// 验签
	err = VerifySignature(params, aksk.Sk, url)
	if err != nil {
		c.Next()
		return
	}
	// 认证通过获取请求来源信息
	server, err := service.ServerSvcIns.GetBySign(aksk.ServerSign)
	if err != nil {
		log.Printf("+%v", err)
		c.Next()
		return
	}
	c.Set("server", server)
	c.Next()
}

func AnalysisHeaderSignature(authorization string) map[string]string {
	var params = make(map[string]string, 5)
	items := strings.Split(authorization, ",")
	for _, item := range items {
		kv := strings.Split(item, "=")
		params[kv[0]] = kv[1]
	}
	return params
}

// VerifySignature
// 示例：Authorization: algorithm=HMAC-SHA256,ak=b4a9bcc1825f11ef,time=1728222161,nonce=29b69dd683e811ef,signature=313233616263746573743132336a6a6b6a3b6c6b6a6b6c3b617364
// 签名规则：sign(key,data) key = sk, data = ak + nonce + time + url
// 验签通过返回ak
func VerifySignature(params map[string]string, sk, url string) error {
	algorithmSign := params["algorithm"]
	algorithm := CheckAlgorithm(algorithmSign)
	if algorithm == nil {
		return errors.New("不支持的算法:" + algorithmSign)
	}
	sign := Signature(algorithm, params["ak"], sk, params["nonce"], params["time"], url)
	if sign == params["signature"] {
		return nil
	}
	return errors.New("验签不通过")
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

// CheckAlgorithm 检测支持的算法
func CheckAlgorithm(algorithm string) func() hash.Hash {
	switch algorithm {
	case "HMAC-SHA256":
		return sha256.New
	case "HMAC-SHA512":
		return sha512.New
	default:
		return nil
	}
}
