package test

import (
	"net/http"
	"testing"
	"bytes"

	"github.com/Piasy/BeegoTDDBootStrap/utils"
)

func TestPostVerificationNoPhone(t *testing.T) {
	request, _ := http.NewRequest("POST", "/v1/verifications/", bytes.NewBuffer([]byte("")))
	checkIsApiErrorRequest(t, "TestPostVerificationNoPhone", request, 403, utils.ERROR_CODE_PARAM_ERROR)
}

func TestPostVerificationInvalidPhone(t *testing.T) {
	request, _ := http.NewRequest("POST", "/v1/verifications/", bytes.NewBuffer([]byte("phone=1880123456")))
	checkIsApiErrorRequest(t, "TestPostVerificationInvalidPhone", request, 403, utils.ERROR_CODE_PARAM_ERROR)
}

func TestPostVerificationSuccess(t *testing.T) {
	verification := createVerification(t, "TestPostVerificationSuccess", "18801234567")
	deleteVerification(t, verification.Id)
}
