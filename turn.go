package gtools

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/axgle/mahonia"
)

var Turn turnInterface

type turnInterface interface {
	StringToInt(value string) int
	IntToString(value int) string
	Int64ToString(value int64) string
	StringToInt64(value string) int64
	ConvertToString(src string, srcCode string, tagCode string) string
	ToString(src interface{}) string
	EncodeByte(data interface{}) ([]byte, error)
	DecodeByte(data []byte, to interface{}) error
	ByteToHex(data []byte) string
	HexToBye(hexStr string) []byte
	UnicodeEmojiDecode(s string) string
	UnicodeEmojiCode(s string) string
	DbcToSbc(str string) string
}

type turn struct{}

func init() {
	Turn = &turn{}
}

// string转成int
func (_turn turn) StringToInt(value string) int {
	if value == "" {
		return 0
	}
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
	if value == "" {
		return 0
	}
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

/**
 * @description: 尝试将一切转换成string
 * @param {interface{}} src
 * @return {*}
 */
func (_turn turn) ToString(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case time.Time:
		return Time.TimeToString(v, "")
	case bool:
		return strconv.FormatBool(v)
	default:
		{
			b, _ := json.Marshal(v)
			return string(b)
		}
	}
	return fmt.Sprintf("%v", src)
}

/**
 * @description: 编码二进制
 * @param {interface{}} data
 * @return {*}
 */
func (_turn turn) EncodeByte(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

/**
 * @description: 解码二进制
 * @param {[]byte} data
 * @param {interface{}} to
 * @return {*}
 */
func (_turn turn) DecodeByte(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}

/**
 * @description: byte转16进制字符串
 * @param {[]byte} data
 * @return {*}
 */
func (_turn turn) ByteToHex(data []byte) string {
	return hex.EncodeToString(data)
}

/**
 * @description: 16进制字符串转[]byte
 * @param {string} hexStr
 * @return {*}
 */
func (_turn turn) HexToBye(hexStr string) []byte {
	hr, _ := hex.DecodeString(hexStr)
	return hr
}

/**
 * @description: Emoji表情解码
 * @param {string} s
 * @return {*}
 */
func (_turn turn) UnicodeEmojiDecode(s string) string {
	//emoji表情的数据表达式
	re := regexp.MustCompile("\\[[\\\\u0-9a-zA-Z]+\\]")
	//提取emoji数据表达式
	reg := regexp.MustCompile("\\[\\\\u|]")
	src := re.FindAllString(s, -1)
	for i := 0; i < len(src); i++ {
		e := reg.ReplaceAllString(src[i], "")
		p, err := strconv.ParseInt(e, 16, 32)
		if err == nil {
			s = strings.Replace(s, src[i], string(rune(p)), -1)
		}
	}
	return s
}

/**
 * @description: Emoji表情转换
 * @param {string} s
 * @return {*}
 */
func (_turn turn) UnicodeEmojiCode(s string) string {
	ret := ""
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if len(string(rs[i])) == 4 {
			u := `[\u` + strconv.FormatInt(int64(rs[i]), 16) + `]`
			ret += u
		} else {
			ret += string(rs[i])
		}
	}
	return ret
}

/**
 * @description: 全角转半角
 * @param {string} str
 * @return {*}
 */
func (_turn turn) DbcToSbc(str string) string {
	numConv := unicode.SpecialCase{
		unicode.CaseRange{
			Lo: 0x3002, // Lo 全角句号
			Hi: 0x3002, // Hi 全角句号
			Delta: [unicode.MaxCase]rune{
				0,               // UpperCase
				0x002e - 0x3002, // LowerCase 转成半角句号
				0,               // TitleCase
			},
		},
		//
		unicode.CaseRange{
			Lo: 0xFF01, // 从全角！
			Hi: 0xFF19, // 到全角 9
			Delta: [unicode.MaxCase]rune{
				0,               // UpperCase
				0x0021 - 0xFF01, // LowerCase 转成半角
				0,               // TitleCase
			},
		},
		unicode.CaseRange{
			Lo: 0xff21, // Lo: 全角 Ａ
			Hi: 0xFF5A, // Hi:到全角 ｚ
			Delta: [unicode.MaxCase]rune{
				0,               // UpperCase
				0x0041 - 0xff21, // LowerCase 转成半角
				0,               // TitleCase
			},
		},
	}

	return strings.ToLowerSpecial(numConv, str)
}
