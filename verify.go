package gtools

import (
	"regexp"
	"strings"
)

var Verify verifyInterface

type verifyInterface interface {
	Email(email string) bool
	Mobile(mobileNum string) bool
	Ip(ip string) bool
	Phone(mobileNum string) bool
	IsMail(username string) (isMail bool)
	IDCard(cardNo string) bool
}

type verify struct{}

func init() {
	Verify = &verify{}
}

func (_verify verify) Email(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func (_verify verify) Mobile(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func (_verify verify) Ip(ip string) bool {
	addr := strings.Trim(ip, " ")
	regStr := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	if match, _ := regexp.MatchString(regStr, addr); match {
		return true
	}
	return false
}

// IsPhone 判断是否是手机号
func (_verify verify) Phone(mobileNum string) bool {
	tmp := `^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\d{8}$`
	reg := regexp.MustCompile(tmp)
	return reg.MatchString(mobileNum)
}

// IsMail 判断用户是否是邮件用户
func (_verify verify) IsMail(username string) (isMail bool) {
	isMail = false
	if strings.Contains(username, "@") {
		isMail = true //是邮箱
	}
	return
}

// IsIDCard 判断是否是18或15位身份证
func (_verify verify) IDCard(cardNo string) bool {
	//18位身份证 ^(\d{17})([0-9]|X)$
	if m, _ := regexp.MatchString(`(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)`, cardNo); !m {
		return false
	}
	return true
}
