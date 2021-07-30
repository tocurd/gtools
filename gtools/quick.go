package gtools

var Quick quickInterface

type quickInterface interface {
	If(condition bool, trueVal, falseVal interface{}) interface{}
}

type quick struct{}

func init() {
	Quick = &quick{}
}

/**
 * @description: 实现三元表达式的功能
 * @param {bool} condition
 * @param {*} trueVal
 * @param {interface{}} falseVal
 * @return {*}
 */
func (_quick quick) If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	} else {
		return falseVal
	}
}
