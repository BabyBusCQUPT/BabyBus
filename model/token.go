package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenClaims struct {
	Identify   string //openId+sessionKey
	Duration   time.Duration
	ExpireTime time.Time
	jwt.StandardClaims
}
