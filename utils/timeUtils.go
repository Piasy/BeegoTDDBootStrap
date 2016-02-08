package utils

import (
	"time"
)

func GetTimeMillis() int64 {
	now := time.Now()
	return now.Unix() * 1000 + int64(now.Nanosecond()) / 1e6
}
