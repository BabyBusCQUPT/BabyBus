package service

import (
	"BabyBus/config"
	"BabyBus/dao"
	"BabyBus/model"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// CreateToken 创建token
func CreateToken(identify string) (string, error) {
	tokenClaims := model.TokenClaims{
		Identify: identify,
		Duration: config.ExpiredDuration,
	}
	tokenClaims.ExpireTime = time.Now().Add(tokenClaims.Duration * time.Minute)
	tokenClaims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: tokenClaims.ExpireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString(config.JwtSecret)
	return tokenString, err
}

// ParseToken 解析token获取token结构体
func ParseToken(tokenString string) (*model.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return config.JwtSecret, nil
		})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*model.TokenClaims)
	if !ok {
		return nil, errors.New("断言失败")
	}
	err = token.Claims.Valid()
	return claims, err
}

// ParseTokenIdentify 解析token.Identify字段获取用户信息
func ParseTokenIdentify(user *model.User, tokenClaims *model.TokenClaims) error {
	WeChatConnection := &model.WeChatConnection{}
	if user.Token == "" {
		return errors.New("token不合理")
	}

	identify := tokenClaims.Identify
	err := json.Unmarshal([]byte(identify), &WeChatConnection)
	if err != nil {
		return err
	}

	user.ID = WeChatConnection.Id

	openId := WeChatConnection.OpenId
	sessionKey := WeChatConnection.SessionKey

	user.OpenId = openId
	user.SessionKey = sessionKey

	return nil
}

func GetIdFromToken(user *model.User) error {
	tokenClaims, err := ParseToken(user.Token)
	if err != nil {
		return errors.New("解析token失败,服务器错误")
	}
	err = ParseTokenIdentify(user, tokenClaims)
	if err != nil {
		return errors.New("获取token内用户信息失败,服务器错误")
	}
	return nil
}

// FindOneWithOpenIdAndSessionKey 通过openId和sessionKey来查询用户token
func FindOneWithOpenIdAndSessionKey(user model.User) (string, error) {
	return dao.FindOneWithOpenidAndSessionKey(user)
}
