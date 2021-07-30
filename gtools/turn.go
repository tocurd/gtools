package gtools

import (
	"strconv"

	"github.com/axgle/mahonia"
)

var Turn turnInterface

type turnInterface interface {
	StringToInt(value string) int
	IntToString(value int) string
	Int64ToString(value int64) string
	StringToInt64(value string) int64
	ConvertToString(src string, srcCode string, tagCode string) string
}

type turn struct{}

func init() {
	Turn = &turn{}
}

// string转成int
func (_turn turn) StringToInt(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		return -1
	}
	return result
}

// int转成string：
func (_turn turn) IntToString(value int) string {
	return strconv.Itoa(value)
}

// int64转成string：
func (_turn turn) Int64ToString(value int64) string {
	return strconv.FormatInt(value, 10)
}

func (_turn turn) StringToInt64(value string) int64 {
	result, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		return -1
	}
	return result
}

/**
 * @description:字符串转码
 * @param {string} src
 * @param {string} srcCode
 * @param {string} tagCode
 * @return {*}
 */
func (_turn turn) ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
