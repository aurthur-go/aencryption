package aencryption

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func Test_AES(t *testing.T) {
	key := []byte("0123456789abcdef")
	str := "hello aurthur"
	result, err := AesEncrypt([]byte(str), key)
	if err != nil {
		t.Error("error")
	}
	fmt.Println("加密后: " + base64.StdEncoding.EncodeToString(result))
	origData, err := AesDecrypt(result, key)
	if err != nil {
		t.Error("error")
	}
	if string(origData) == str {
		t.Log("right")
	}
}
