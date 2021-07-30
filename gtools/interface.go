package gtools

import (
	"errors"
	"fmt"
	"reflect"
)

var Interface interfaceInterface

type interfaceInterface interface {
	ToString(v interface{}) (result string, err error)
}

type selfInterface struct{}

func init() {
	Interface = &selfInterface{}
}

func (_interface selfInterface) ToString(v interface{}) (result string, err error) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		result = fmt.Sprintf("%v", v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result = fmt.Sprintf("%v", v)
	case reflect.String:
		result = v.(string)
	default:
		err = errors.New("can't transition to string")
	}
	return result, err
}
