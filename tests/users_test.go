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
	checkIsApiErrorRequest(t, "TestGetUserParamsErrorNoUid", request, 401, utils.ERROR_CODE_TOKENS_INVALID_TOKEN)
}

func TestGetUserParamsErrorNoToken(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/users/2971788563", nil)
	checkIsApiErrorRequest(t, "TestGetUserParamsErrorNoToken", request, 401, utils.ERROR_CODE_TOKENS_INVALID_TOKEN)
}

func TestGetUserInvalidToken(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/users/2971788563?token=lgJYnQXrKVPoInPTPnokdPOZISzosxQzNUceRJyA", nil)
	checkIsApiErrorRequest(t, "TestGetUserInvalidToken", request, 401, utils.ERROR_CODE_TOKENS_INVALID_TOKEN)
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

func TestPostUserBasicAuthFail(t *testing.T) {
	request, _ := http.NewRequest("POST", "/v1/users/",
		bytes.NewBuffer([]byte("phone=1881234567&secret=8428d916f8cca9ba5971bf58b34d38da20bc3dff&code=123456")))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	checkIsApiErrorRequest(t, "TestPostUserInvalidPhone", request, 401, utils.ERROR_CODE_BASIC_AUTH_FAIL)
}

func TestPostUserInvalidPhone(t *testing.T) {
	request, _ := http.NewRequest("POST", "/v1/users/",
		bytes.NewBuffer([]byte("phone=1881234567&secret=8428d916f8cca9ba5971bf58b34d38da20bc3dff&code=123456")))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", "dGVzdF9jbGllbnQ6dGVzdF9wYXNz")
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
	request.Header.Add("Authorization", "dGVzdF9jbGllbnQ6dGVzdF9wYXNz")
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
	request.Header.Add("Authorization", "dGVzdF9jbGllbnQ6dGVzdF9wYXNz")
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

	request, _ := http.NewRequest("POST", "/v1/verifications/", bytes.NewBuffer([]byte("phone=" + phone)))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", "dGVzdF9jbGllbnQ6dGVzdF9wYXNz")
	checkIsApiErrorRequest(t, "TestPostUserExists", request, 422, utils.ERROR_CODE_USERS_PHONE_REGISTERED)

	deleteVerification(t, verification.Id)
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

func TestPostTwoUserSuccess(t *testing.T) {
	phone := "18801234567"
	secret := "8428d916f8cca9ba5971bf58b34d38da20bc3dff"
	phone2 := "18801234568"

	verification := createVerification(t, "TestPostTwoUserSuccess(first)", phone)
	user := createUser(t, "TestPostTwoUserSuccess(first)", phone, secret, verification.Code)
	checkHasUserByUid(t, "TestPostTwoUserSuccess(first)", user)

	verification2 := createVerification(t, "TestPostTwoUserSuccess(second)", phone2)
	user2 := createUser(t, "TestPostTwoUserSuccess(second)", phone2, secret, verification2.Code)
	checkHasUserByUid(t, "TestPostTwoUserSuccess(second)", user2)

	deleteVerification(t, verification.Id)
	deleteUser(t, user.Id)
	deleteVerification(t, verification2.Id)
	deleteUser(t, user2.Id)
}

func TestPatchUserInfoNoToken(t *testing.T) {
	phone := "18801234567"
	secret := "8428d916f8cca9ba5971bf58b34d38da20bc3dff"

	verification := createVerification(t, "TestPatchUserInfoNoToken", phone)
	user := createUser(t, "TestPatchUserInfoNoToken", phone, secret, verification.Code)
	checkHasUserByUid(t, "TestPatchUserInfoNoToken", user)

	request, _ := http.NewRequest("PATCH", "/v1/users/", bytes.NewBuffer([]byte("")))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	checkIsApiErrorRequest(t, "TestPatchUserInfoNoToken", request, 401, utils.ERROR_CODE_TOKENS_INVALID_TOKEN)

	deleteVerification(t, verification.Id)
	deleteUser(t, user.Id)
}

// TODO test update praises
func TestPatchUserInfoOneByOne(t *testing.T) {
	phone := "18801234567"
	secret := "8428d916f8cca9ba5971bf58b34d38da20bc3dff"

	verification := createVerification(t, "TestPatchUserInfoOneByOne", phone)
	user := createUser(t, "TestPatchUserInfoOneByOne", phone, secret, verification.Code)

	updated := patchUserInfo(t, "TestPatchUserInfoOneByOne", user.Token,
		[]byte(fmt.Sprintf("token=%s&nickname=%s", user.Token, "Piasy")))
	user.Nickname = updated.Nickname
	user.UpdateAt = updated.UpdateAt
	checkHasUserByUid(t, "TestPatchUserInfoOneByOne", user)

	updated = patchUserInfo(t, "TestPatchUserInfoOneByOne", user.Token,
		[]byte(fmt.Sprintf("token=%s&gender=%s", user.Token, "1")))
	user.Gender = updated.Gender
	user.UpdateAt = updated.UpdateAt
	checkHasUserByUid(t, "TestPatchUserInfoOneByOne", user)

	updated = patchUserInfo(t, "TestPatchUserInfoOneByOne", user.Token,
		[]byte(fmt.Sprintf("token=%s&avatar=%s", user.Token,
			"https://avatars2.githubusercontent.com/u/3098704?v=3&s=460")))
	user.Avatar = updated.Avatar
	user.UpdateAt = updated.UpdateAt
	checkHasUserByUid(t, "TestPatchUserInfoOneByOne", user)

	phone2 := "18801234568"
	verification2 := createVerification(t, "TestPatchUserInfoOneByOne", phone2)
	updated = patchUserInfo(t, "TestPatchUserInfoOneByOne", user.Token,
		[]byte(fmt.Sprintf("token=%s&phone=%s&code=%s", user.Token, phone2, verification2.Code)))
	user.Phone = updated.Phone
	user.UpdateAt = updated.UpdateAt
	checkHasUserByUid(t, "TestPatchUserInfoOneByOne", user)

	// other credential test is omitted...

	deleteVerification(t, verification.Id)
	deleteVerification(t, verification2.Id)
	deleteUser(t, user.Id)
}

func TestPatchUserInfoAll(t *testing.T) {
	phone := "18801234567"
	secret := "8428d916f8cca9ba5971bf58b34d38da20bc3dff"
	phone2 := "18801234568"

	verification := createVerification(t, "TestPatchUserInfoOneByOne", phone)
	user := createUser(t, "TestPatchUserInfoOneByOne", phone, secret, verification.Code)
	checkHasUserByUid(t, "TestPatchUserInfoOneByOne", user)

	verification2 := createVerification(t, "TestPatchUserInfoOneByOne", phone2)
	updated := patchUserInfo(t, "TestPatchUserInfoOneByOne", user.Token,
		[]byte(fmt.Sprintf("token=%s&phone=%s&code=%s&nickname=%s&description=%s&gender=%s&school_id=%s&major_id=%s&birthday=%s&avatar=%s",
			user.Token, phone2, verification2.Code, "Piasy", "I'm Piasy", "1", "1001", "5", "1447166995",
			"https://avatars2.githubusercontent.com/u/3098704?v=3&s=460")))
	checkHasUserByUid(t, "TestPatchUserInfoOneByOne", updated)

	deleteVerification(t, verification.Id)
	deleteUser(t, user.Id)
}

func TestPatchUserInfoOneByOneWithIllegalParams(t *testing.T) {
	phone := "18801234567"
	secret := "8428d916f8cca9ba5971bf58b34d38da20bc3dff"

	verification := createVerification(t, "TestPatchUserInfoOneByOneWithIllegalParams", phone)
	user := createUser(t, "TestPatchUserInfoOneByOneWithIllegalParams", phone, secret, verification.Code)
	checkHasUserByUid(t, "TestPatchUserInfoOneByOneWithIllegalParams", user)

	request, _ := http.NewRequest("PATCH", "/v1/users/",
		bytes.NewBuffer([]byte(fmt.Sprintf("token=%s&nickname=%s", user.Token, "PiasyPiasyPiasyPiasyPiasyPiasy"))))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	checkIsApiErrorRequest(t, "TestPatchUserInfoOneByOneWithIllegalParams", request, 403, utils.ERROR_CODE_USERS_INVALID_NICKNAME)

	request, _ = http.NewRequest("PATCH", "/v1/users/",
		bytes.NewBuffer([]byte(fmt.Sprintf("token=%s&gender=%s", user.Token, "3"))))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	checkIsApiErrorRequest(t, "TestPatchUserInfoOneByOneWithIllegalParams", request, 403, utils.ERROR_CODE_USERS_INVALID_GENDER_VALUE)

	deleteVerification(t, verification.Id)
	deleteUser(t, user.Id)
}
