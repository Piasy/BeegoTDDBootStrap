package utils

import (
	"time"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const USER_MIN_UID int64 = 1000000000
const VERIFY_CODE_LEN int = 6

const VERIFY_CODE_EXPIRE_IN_SECONDS int64 = 10 * 60

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1 << letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var mSrc = rand.NewSource(time.Now().UnixNano())
var mRand = rand.New(mSrc)
