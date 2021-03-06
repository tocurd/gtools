package gtools

import (
	"fmt"
	"reflect"
)

var Interface interfaceInterface

type interfaceInterface interface {
	ToString(v interface{}) (result string)
}

type selfInterface struct{}

func init() {
	Interface = &selfInterface{}
}

// string转成int：
// int, err := strconv.Atoi(string)
// string转成int64：
// int64, err := strconv.ParseInt(string, 10, 64)
// int转成string：
// string := strconv.Itoa(int)
// int64转成string：
// string := strconv.FormatInt(int64,10)
func (_interface selfInterface) ToString(v interface{}) (result string) {
	if v == nil {
		return ""
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Int64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		result = fmt.Sprintf("%v", v)
	case reflect.Float32, reflect.Float64:
		result = fmt.Sprintf("%v", v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result = fmt.Sprintf("%v", v)
	case reflect.String:
		result = v.(string)
	default:
		return ""
	}
	return result
}
