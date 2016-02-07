package utils

type ApiError struct {
	Code    int `json:"code"`
	Message string `json:"message"`
	Request string `json:"request"`
}

func Issue(code int, request string) ApiError {
	return ApiError{code, ERROR_MESSAGES[code], request}
}

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

// verify code
	ERROR_CODE_VERIFY_CODE_MISMATCH = 20201

// tokens
	ERROR_CODE_TOKENS_PASSWORD_MISMATCH = 20301
	ERROR_CODE_TOKENS_INVALID_TOKEN = 20302
)

var (
	ERROR_MESSAGES = map[int]string{
		// system level
		ERROR_CODE_SYSTEM_ERROR: "Oops, System error",
		ERROR_CODE_PARAM_ERROR: "Param error, see doc for more info",

		// users
		ERROR_CODE_USERS_USER_NOT_EXISTS: "Queried user not exists",
		ERROR_CODE_USERS_PHONE_REGISTERED: "Phone has registered",

		// verify code
		ERROR_CODE_VERIFY_CODE_MISMATCH: "Verify code mismatch",

		// tokens
		ERROR_CODE_TOKENS_PASSWORD_MISMATCH: "Password mismatch",
		ERROR_CODE_TOKENS_INVALID_TOKEN: "Invalid token",
	}
)
