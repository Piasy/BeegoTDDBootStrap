package utils

import (
	"time"
	"math/rand"
)

type ApiError struct {
	Code    int `json:"code"`
	Message string `json:"message"`
	Request string `json:"request"`
}

func Issue(code int, request string) ApiError {
	return ApiError{code, ERROR_MESSAGES[code], request}
}

const (
	GENDER_UNKNOWN = 0
	GENDER_MALE = 1 << iota
	GENDER_FEMALE
)

const (
	SNS_PLATFORM_WEIXIN = iota
	SNS_PLATFORM_WEIBO
	SNS_PLATFORM_QQ
)

const USER_NICKNAME_MEX_LEN int = 20
const USER_AVATAR_MEX_LEN int = 255
const USER_MIN_UID int64 = 1000000000
const VERIFY_CODE_LEN int = 6
const VERIFY_CODE_EXPIRE_IN_SECONDS int64 = 10 * 60

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1 << letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var mSrc = rand.NewSource(time.Now().UnixNano())
var mRand = rand.New(mSrc)


// error code definition, three part, [level][function][number]
// [level], 1: system level, 2: service level
// [function] 01: tokens; 02: verify code; ...
// [number] detailed error number
const (
// system level
	ERROR_CODE_SYSTEM_ERROR = 10001
	ERROR_CODE_PARAM_ERROR = 10002

// users
	ERROR_CODE_USERS_USER_NOT_EXISTS = 20101
	ERROR_CODE_USERS_PHONE_REGISTERED = 20102
	ERROR_CODE_USERS_INVALID_NICKNAME = 20103
	ERROR_CODE_USERS_INVALID_GENDER_VALUE = 20104
	ERROR_CODE_USERS_INVALID_AVATAR = 20105

// verify code
	ERROR_CODE_VERIFY_CODE_MISMATCH = 20201

// tokens
	ERROR_CODE_TOKENS_PASSWORD_MISMATCH = 20301
	ERROR_CODE_TOKENS_INVALID_TOKEN = 20302

// auth
	ERROR_CODE_AUTH_WEIXIN_AUTH_FAIL = 20401
	ERROR_CODE_AUTH_WEIBO_AUTH_FAIL = 20402
	ERROR_CODE_AUTH_QQ_AUTH_FAIL = 20403

// basic auth
	ERROR_CODE_BASIC_AUTH_FAIL = 20801

)

var (
	ERROR_MESSAGES = map[int]string{
		// system level
		ERROR_CODE_SYSTEM_ERROR: "Oops, System error",
		ERROR_CODE_PARAM_ERROR: "Param error, see doc for more info",

		// users
		ERROR_CODE_USERS_USER_NOT_EXISTS: "Queried user not exists",
		ERROR_CODE_USERS_PHONE_REGISTERED: "Phone has registered",
		ERROR_CODE_USERS_INVALID_NICKNAME: "Nickname invalid",
		ERROR_CODE_USERS_INVALID_GENDER_VALUE: "Invalid gender value, should only be 1 or 2",
		ERROR_CODE_USERS_INVALID_AVATAR: "Avatar url invalid",

		// verify code
		ERROR_CODE_VERIFY_CODE_MISMATCH: "Verify code mismatch",

		// tokens
		ERROR_CODE_TOKENS_PASSWORD_MISMATCH: "Password mismatch",
		ERROR_CODE_TOKENS_INVALID_TOKEN: "Invalid token",

		// auth
		ERROR_CODE_AUTH_WEIXIN_AUTH_FAIL: "Weixin auth fail",
		ERROR_CODE_AUTH_WEIBO_AUTH_FAIL: "Weibo auth fail",
		ERROR_CODE_AUTH_QQ_AUTH_FAIL: "QQ auth fail",

		// basic auth
		ERROR_CODE_BASIC_AUTH_FAIL: "Basic auth fail",

	}
)
