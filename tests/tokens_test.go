package test

import (
	"net/http"
	"testing"
	"bytes"
	"fmt"

	"github.com/Piasy/BeegoTDDBootStrap/utils"
)

func TestPostTokenParamsError(t *testing.T) {
	request, _ := http.NewRequest("POST", "/v1/tokens/", bytes.NewBuffer([]byte("")))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	checkIsApiErrorRequest(t, "TestPostTokenParamsError", request, 403, utils.ERROR_CODE_PARAM_ERROR)
}

func TestPostTokenPhoneNotExists(t *testing.T) {
	request, _ := http.NewRequest("POST", "/v1/tokens/",
		bytes.NewBuffer([]byte("phone=18812345678&secret=8428d916f8cca9ba5971bf58b34d38da20bc3dff")))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	checkIsApiErrorRequest(t, "TestPostTokenPhoneNotExists", request, 422, utils.ERROR_CODE_USERS_USER_NOT_EXISTS)
}

func TestPostTokenInvalidPhone(t *testing.T) {
	request, _ := http.NewRequest("POST", "/v1/tokens/",
		bytes.NewBuffer([]byte("phone=1881234567&secret=8428d916f8cca9ba5971bf58b34d38da20bc3dff")))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	checkIsApiErrorRequest(t, "TestPostTokenInvalidPhone", request, 403, utils.ERROR_CODE_PARAM_ERROR)
}

func TestPostTokenWrongSecret(t *testing.T) {
	phone := "18801234567"
	secret := "8428d916f8cca9ba5971bf58b34d38da20bc3dff"

	verification := createVerification(t, "TestPostTokenWrongSecret", phone)
	user := createUser(t, "TestPostTokenWrongSecret", phone, secret, verification.Code)

	request, _ := http.NewRequest("POST", "/v1/tokens/",
		bytes.NewBuffer([]byte(fmt.Sprintf("phone=%s&secret=%s", phone, "a" + secret[1:]))))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	checkIsApiErrorRequest(t, "TestPostTokenWrongSecret", request, 422, utils.ERROR_CODE_TOKENS_PASSWORD_MISMATCH)

	deleteVerification(t, verification.Id)
	deleteUser(t, user.Id)
}

func TestPostTokenSuccess(t *testing.T) {
	phone := "18801234567"
	secret := "8428d916f8cca9ba5971bf58b34d38da20bc3dff"

	verification := createVerification(t, "TestPostUserSuccess", phone)
	user := createUser(t, "TestPostTokenSuccess", phone, secret, verification.Code)
	updateTokenByPhone(t, "TestPostTokenSuccess", user, secret)

	deleteVerification(t, verification.Id)
	deleteUser(t, user.Id)
}
