package config

import (
	"time"
)

var JWTSecret = []byte("!tickstatsoulter.,.")

func JWTExpirationDuration() time.Duration {
	return time.Hour * 24 * 7 // 1 week
}
