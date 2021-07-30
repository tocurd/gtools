package gtools

import "strings"

var String stringInterface

type stringInterface interface {
	GetBetweenStr(str, start, end string) string
	GetRight(str string, left string) string
}

type selfString struct{}

func init() {
	String = &selfString{}
}
func (_selfString selfString) GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	}
	str = string([]byte(str)[n+len(start):])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}

func (_selfString selfString) GetRight(str string, left string) string {
	leftIndex := strings.Index(str, left)
	return str[leftIndex+len(left):]
}
