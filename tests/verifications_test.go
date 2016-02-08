package test

import (
	"net/http"
	"testing"
	"bytes"

	"github.com/Piasy/BeegoTDDBootStrap/utils"
)

func TestPostVerificationBasicAuthFail(t *testing.T) {
	request, _ := http.NewRequest("POST", "/v1/verifications/", bytes.NewBuffer([]byte("")))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	checkIsApiErrorRequest(t, "TestPostVerificationNoPhone", request, 401, utils.ERROR_CODE_BASIC_AUTH_FAIL)
}

func TestPostVerificationNoPhone(t *testing.T) {
	request, _ := http.NewRequest("POST", "/v1/verifications/", bytes.NewBuffer([]byte("")))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", "dGVzdF9jbGllbnQ6dGVzdF9wYXNz")
	checkIsApiErrorRequest(t, "TestPostVerificationNoPhone", request, 403, utils.ERROR_CODE_PARAM_ERROR)
}

func TestPostVerificationInvalidPhone(t *testing.T) {
	request, _ := http.NewRequest("POST", "/v1/verifications/", bytes.NewBuffer([]byte("phone=1880123456")))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", "dGVzdF9jbGllbnQ6dGVzdF9wYXNz")
	checkIsApiErrorRequest(t, "TestPostVerificationInvalidPhone", request, 403, utils.ERROR_CODE_PARAM_ERROR)
}

func TestPostVerificationSuccess(t *testing.T) {
	verification := createVerification(t, "TestPostVerificationSuccess", "18801234567")
	deleteVerification(t, verification.Id)
}
