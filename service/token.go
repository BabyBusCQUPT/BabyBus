package service

import (
	"BabyBus/config"
	"BabyBus/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

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
