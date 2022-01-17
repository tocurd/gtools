package gtools

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtInterface interface {
}

type Jwt struct {
	JwtSecret []byte
}

//Claim是一些实体（通常指的用户）的状态和额外的元数据
type jwtClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

/**
 * @description: 根据用户的用户名产生token
 * @param {string} username
 * @return {*}
 */
func (_selfJwt Jwt) Generate(username string, expire time.Duration) (string, error) {
	//设置token有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(expire)

	claims := jwtClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(_selfJwt.JwtSecret)
	return token, err
}

/**
 * @description: 根据传入的token值获取到对象信息
 * @param {string} token
 * @return {*}
 */
func (_selfJwt Jwt) Parse(token string) (*jwtClaims, error) {

	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return _selfJwt.JwtSecret, nil
	})

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*jwtClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}
