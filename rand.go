package gtools

import (
	"math/rand"
	"time"
)

var Rand randInterface

type randInterface interface {
	Random(length int, char []byte) string
	RangeNumString(length int) string
	RangeNum(length int) int
	RandInt(min int, max int) int
}

type selfRand struct{}

func init() {
	Rand = &selfRand{}
}

//生成随机字符串
var r *rand.Rand

/**
 * @description: 生成随机字符串
 * @param {int} n 长度
 * @param {[]byte} char 自定义的数据
 * @return {*}
 */
func (_selfRand selfRand) Random(length int, char []byte) string {
	if len(char) <= 0 {
		char = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
	result := []byte{}
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	for i := 0; i < length; i++ {
		result = append(result, char[r.Intn(len(char))])
	}
	return string(result)
}

/**
 * @description: 生成随机数字字符串
 * @param {int} n
 * @return {*}
 */
func (_selfRand selfRand) RangeNumString(length int) string {
	var _bytes = []byte("0123456789")
	var r *rand.Rand

	result := []byte{}
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	for i := 0; i < length; i++ {
		result = append(result, _bytes[r.Intn(len(_bytes))])
	}
	return string(result)
}

/**
 * @description: 生成随机整数
 * @param {int} length 位数
 * @return {*}
 */
func (_selfRand selfRand) RangeNum(length int) int {
	var max, min int = 1, 1
	if length > 0 {
		for i := 0; i < length; i++ {
			max = max * 10
		}
		for i := 0; i < length-1; i++ {
			min = min * 10
		}
	}
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

/**
 * @description: 生成指定范围的随机数
 * @param {int} min
 * @param {int} max
 * @return {*}
 */
func (_selfRand selfRand) RandInt(min int, max int) int {
	if min > max {
		min = 0
		max = 0
	}
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
