package util

import (
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

func Test1(t *testing.T) {
	t.Log(hex.EncodeToString([]byte("123abctest123jjkj;lkjkl;asd")))
}

func Test2(t *testing.T) {
	t.Log(time.Now().Unix())
}
