package gtools

import (
	cryptoMd5 "crypto/md5"
	"fmt"
)

var Hash hashInterface

type hashInterface interface {
	Md5(value string) string
}

type hash struct{}

func init() {
	Hash = &hash{}
}

func (_hash hash) Md5(value string) string {
	return fmt.Sprintf("%x", cryptoMd5.Sum([]byte(value)))
}
