package utils

import (
	"fmt"
)

func randStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n - 1, mSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = mSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// 40 characters random token
func GenToken() string {
	return randStringBytesMaskImprSrc(40)
}

func GenUid() int64 {
	return int64(mRand.Uint32()) + USER_MIN_UID
}

func GenVerifyCode() string {
	return fmt.Sprintf("%06d", mRand.Intn(100000))
}
