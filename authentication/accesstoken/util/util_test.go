package util

import "testing"

func TestGenerate16Str(t *testing.T) {
	str := Generate16Str()
	t.Log(str)
}

func TestGenerate24Str(t *testing.T) {
	str := Generate24Str()
	t.Log(str)
}
