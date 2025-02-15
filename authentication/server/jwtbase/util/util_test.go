package util

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
	"time"
)

func TestGenerate16Str(t *testing.T) {
	str := Generate16Str()
	t.Log(str)
}

func TestGenerate24Str(t *testing.T) {
	str := Generate24Str()
	t.Log(str)
}

func TestGenerate32Str(t *testing.T) {
	str := GenerateStr()
	t.Log(str)
}

func Test1(t *testing.T) {
	t.Log(hex.EncodeToString([]byte("123abctest123jjkj;lkjkl;asd")))
}

func Test2(t *testing.T) {
	t.Log(time.Now().Unix())
}

func TestSignature(t *testing.T) {
	data, key := []byte("123abc"), []byte("123abc")
	sign, err := Signature(data, key, "HS512")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(base64.RawURLEncoding.EncodeToString(sign))
}
