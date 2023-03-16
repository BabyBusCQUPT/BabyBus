package config

import "time"

var (
	JwtSecret = []byte("BabyBus.com")
)

const (
	ExpiredDuration = time.Minute * 1
)
