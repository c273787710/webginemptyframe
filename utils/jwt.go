package utils

import (
	"adminframe/framework/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTInfo struct {
	UID int `json:"uid"`
	Username string `json:"username"`
	ClientIP string `json:"client_ip"`
	UserAgent string `json:"user_agent"`
	jwt.StandardClaims
}
type UserTemplate struct {
	ID int
	Username string
	ClientIP string
	UserAgent string
}
//生成jwt加密字符串
func GenerateJWTToken(u UserTemplate)(string,error){
	nowTime := time.Now()
	expireTime := nowTime.Add( config.JWTSetting.ExpireTime)
	jwtInfo := &JWTInfo{}
	jwtInfo.UID = u.ID
	jwtInfo.Username = u.Username
	jwtInfo.ClientIP = u.ClientIP
	jwtInfo.UserAgent = u.UserAgent
	jwtInfo.ExpiresAt = expireTime.Unix()
	jwtInfo.Issuer = config.JWTSetting.Issuer
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwtInfo)
	jwtSecret := []byte(config.JWTSetting.Secret)
	return token.SignedString(jwtSecret)
}

//解析jwt
func ParseJWTToken(token string)(*UserTemplate,error){
	jwtSecret := []byte(config.JWTSetting.Secret)
	tokenClaims,err := jwt.ParseWithClaims(token,&JWTInfo{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret,nil
	})
	if tokenClaims != nil {
		if jwtInfo,ok := tokenClaims.Claims.(*JWTInfo);ok && tokenClaims.Valid{
			userTemplate := &UserTemplate{
				ID: jwtInfo.UID,
				Username: jwtInfo.Username,
				ClientIP: jwtInfo.ClientIP,
				UserAgent: jwtInfo.UserAgent,
			}
			return userTemplate,nil
		}
	}
	return nil,err
}