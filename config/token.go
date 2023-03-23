package config

import "time"

var (
	JwtSecret    = []byte("BabyBus.com")
	InvalidToken = ""
)

const (
	ExpiredDuration = time.Minute * 1
)
