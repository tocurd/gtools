package gtools

import (
	"errors"
	"reflect"
	"strings"
)

var Array arrayInterface

type arrayInterface interface {
	Find(target interface{}, value interface{}) bool
	Join(a []string, sep string) string
	Field(array interface{}, key ...string) (result []interface{}, err error)
	Column(array interface{}, key string) (result []interface{}, err error)
}

type array struct{}

func init() {
	Array = &array{}
}

/**
 * @description: 查找某个值是否在数组内
 * @param {interface{}} target
 * @param {interface{}} value
 * @return {*}
 */
func (_array array) Find(target interface{}, value interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == value {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(value)).IsValid() {
			return true
		}
	}
	return false
}

/**
 * @description: 数组转字符串
 * @param {[]string} a
 * @param {string} sep
 * @return {*}
 */
func (_array array) Join(a []string, sep string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0]
	case 2:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return a[0] + sep + a[1]
	case 3:
		// Special case for common small values.
		// Remove if golang.org/issue/6714 is fixed
		return a[0] + sep + a[1] + sep + a[2]
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}

/**
 * @description: 获取数组的指定字段
 * @param {interface{}} array
 * @param {...string} key
 * @return {*}
 */
func (_array array) Field(array interface{}, key ...string) (result []interface{}, err error) {
	result = make([]interface{}, 0)
	typeof := reflect.TypeOf(array)
	value := reflect.ValueOf(array)
	if typeof.Kind() != reflect.Slice {
		return nil, errors.New("array type not slice")
	}
	if value.Len() == 0 {
		return result, nil
	}

	for i := 0; i < value.Len(); i++ {
		indexv := value.Index(i)

		// 判断数据是否是一个结构体
		kind := indexv.Type().Kind()
		if kind != reflect.Struct {
			return nil, errors.New("element type not struct")
		}

		row := make(map[string]interface{}, 1)
		for keyIndex := 0; keyIndex < len(key); keyIndex++ {
			keyValue := key[keyIndex]
			keyTurnValue := ""

			// 是否有需要转换的名字
			if strings.Contains(keyValue, "->") {
				split := strings.Split(keyValue, "->")
				if len(split) < 2 {
					return nil, errors.New("split field key error")
				}
				keyValue = split[0]
				keyTurnValue = split[1]
			}

			// 在这里开始获取值
			mapKeyInterface := indexv.FieldByName(keyValue)
			if mapKeyInterface.Kind() == reflect.Invalid {
				return nil, errors.New("key not exist")
			}
			resultKeyValue := mapKeyInterface.Interface()

			if keyTurnValue == "" {
				row[keyValue] = resultKeyValue
			} else {
				row[keyTurnValue] = resultKeyValue
			}
		}

		result = append(result, row)
	}
	return result, nil
}

/**
 * @description: 将某个字段栏目里的数据提取
 * @param {interface{}} array
 * @param {string} key
 * @return {*}
 */
func (_array array) Column(array interface{}, key string) (result []interface{}, err error) {
	result = make([]interface{}, 0)
	t := reflect.TypeOf(array)
	v := reflect.ValueOf(array)
	if t.Kind() != reflect.Slice {
		return nil, errors.New("array type not slice")
	}
	if v.Len() == 0 {
		return nil, errors.New("array len is zero")
	}

	for i := 0; i < v.Len(); i++ {
		indexv := v.Index(i)

		if indexv.Type().Kind() != reflect.Struct {
			return nil, errors.New("element type not struct")
		}
		mapKeyInterface := indexv.FieldByName(key)
		if mapKeyInterface.Kind() == reflect.Invalid {
			return nil, errors.New("key not exist")
		}
		mapKeyString := Interface.ToString(mapKeyInterface.Interface())

		result = append(result, mapKeyString)
	}
	return result, err
}
