package gtools

import (
	encodingBase64 "encoding/base64"
)

var Base64 base64Interface

type base64Interface interface {
	Encode(src []byte) string
	Decode(src string) ([]byte, error)
}

type base64 struct{}

func init() {
	Base64 = &base64{}
}

// Encode base64 编码
func (_base64 base64) Encode(src []byte) string {
	return encodingBase64.StdEncoding.EncodeToString(src)
}

// Decode base64 解码
func (_base64 base64) Decode(src string) ([]byte, error) {
	return encodingBase64.StdEncoding.DecodeString(src)
}
