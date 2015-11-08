package test

import (
	"fmt"
	"net/http"
	"testing"
	"bytes"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/stretchr/testify/assert"

	"github.com/Piasy/BeegoTDDBootStrap/utils"
)

func TestGetUserParamsErrorNoUid(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/users/", nil)
	checkIsApiErrorRequest(t, "TestGetUserParamsErrorNoUid", request, 403, utils.ERROR_CODE_PARAM_ERROR)
}

func TestGetUserParamsErrorNoToken(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/users/2971788563", nil)
	checkIsApiErrorRequest(t, "TestGetUserParamsErrorNoToken", request, 403, utils.ERROR_CODE_PARAM_ERROR)
}

func TestGetUserInvalidToken(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/users/2971788563?token=lgJYnQXrKVPoInPTPnokdPOZISzosxQzNUceRJyA", nil)
	checkIsApiErrorRequest(t, "TestGetUserInvalidToken", request, 403, utils.ERROR_CODE_TOKENS_INVALID_TOKEN)
}

func TestGetUserNotExists(t *testing.T) {
	phone := "18801234567"
	secret := "8428d916f8cca9ba5971bf58b34d38da20bc3dff"

	verification := createVerification(t, "TestGetUserNotExists", phone)
	user := createUser(t, "TestGetUserNotExists", phone, secret, verification.Code)
	request, _ := http.NewRequest("GET", fmt.Sprintf("/v1/users/%d?token=%s", user.Uid + 1, user.Token), nil)
	checkIsApiErrorRequest(t, "TestGetUserNotExists", request, 404, utils.ERROR_CODE_USERS_USER_NOT_EXISTS)

	deleteVerification(t, verification.Id)
	deleteUser(t, user.Id)
}

func TestPostUserInvalidPhone(t *testing.T) {
		request, _ := http.NewRequest("POST", "/v1/users/",
			bytes.NewBuffer([]byte("phone=1881234567&secret=8428d916f8cca9ba5971bf58b34d38da20bc3dff&code=123456")))
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		checkIsApiErrorRequest(t, "TestPostUserInvalidPhone", request, 403, utils.ERROR_CODE_PARAM_ERROR)
}

func TestPostUserExpiredCode(t *testing.T) {
	phone := "18801234567"
	secret := "8428d916f8cca9ba5971bf58b34d38da20bc3dff"

	verification := createVerification(t, "TestPostUserExpiredCode", phone)
	// simulate expire
	o := orm.NewOrm()
	verification.Expire = time.Now().Unix() - 100
	_, err := o.Update(verification)
	assert.Nil(t, err)

	request, _ := http.NewRequest("POST", "/v1/users/",
		bytes.NewBuffer([]byte(fmt.Sprintf("phone=%s&secret=%s&code=%s", phone, secret, verification.Code))))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	checkIsApiErrorRequest(t, "TestPostUserExpiredCode", request, 422, utils.ERROR_CODE_VERIFY_CODE_MISMATCH)

	deleteVerification(t, verification.Id)
}

func TestPostUserDupCode(t *testing.T) {
	phone := "18801234567"
	secret := "8428d916f8cca9ba5971bf58b34d38da20bc3dff"

	verification := createVerification(t, "TestPostUserDupCode", phone)
	user := createUser(t, "TestPostUserDupCode", phone, secret, verification.Code)
	checkHasUserByUid(t, "TestPostUserDupCode", user)

	phone2 := "18801234568"
	request, _ := http.NewRequest("POST", "/v1/users/",
		bytes.NewBuffer([]byte(fmt.Sprintf("phone=%s&secret=%s&code=%s", phone2, secret, verification.Code))))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	checkIsApiErrorRequest(t, "TestPostUserDupCode", request, 422, utils.ERROR_CODE_VERIFY_CODE_MISMATCH)

	deleteVerification(t, verification.Id)
	deleteUser(t, user.Id)
}

func TestPostUserExists(t *testing.T) {
	phone := "18801234567"
	secret := "8428d916f8cca9ba5971bf58b34d38da20bc3dff"

	verification := createVerification(t, "TestPostUserSuccess", phone)
	user := createUser(t, "TestPostUserExists", phone, secret, verification.Code)
	checkHasUserByUid(t, "TestPostUserExists", user)

	verification2 := createVerification(t, "TestPostUserExists", phone)
	request, _ := http.NewRequest("POST", "/v1/users/",
		bytes.NewBuffer([]byte(fmt.Sprintf("phone=%s&secret=%s&code=%s", phone, secret, verification2.Code))))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	checkIsApiErrorRequest(t, "TestPostUserExists", request, 422, utils.ERROR_CODE_USERS_PHONE_REGISTERED)

	deleteVerification(t, verification.Id)
	deleteVerification(t, verification2.Id)
	deleteUser(t, user.Id)
}

func TestPostUserSuccess(t *testing.T) {
	phone := "18801234567"
	secret := "8428d916f8cca9ba5971bf58b34d38da20bc3dff"

	verification := createVerification(t, "TestPostUserSuccess", phone)
	user := createUser(t, "TestPostUserSuccess", phone, secret, verification.Code)
	checkHasUserByUid(t, "TestPostUserSuccess", user)

	deleteVerification(t, verification.Id)
	deleteUser(t, user.Id)
}
