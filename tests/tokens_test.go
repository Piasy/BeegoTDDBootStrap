package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"fmt"
	"encoding/json"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/Piasy/BeegoTDDBootStrap/utils"
	"github.com/Piasy/BeegoTDDBootStrap/models"
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

func updateTokenByPhone(t *testing.T, subject string, expect *models.User, secret string) {
	request, _ := http.NewRequest("POST", "/v1/tokens/",
		bytes.NewBuffer([]byte(fmt.Sprintf("phone=%s&secret=%s", *expect.Phone, secret))))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", "dGVzdF9jbGllbnQ6dGVzdF9wYXNz")
	recorder := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(recorder, request)
	beego.Debug("testing <", subject, ">: verify user, Code[", recorder.Code, "]\n", recorder.Body.String())

	Convey("Subject: " + subject + ": verify user\n", t, func() {
		Convey("Get user status code should be 201", func() {
			soResponseWithStatusCode(recorder, 201)
		})
		Convey("Get user should be same as created", func() {
			var got models.User
			err := json.Unmarshal(recorder.Body.Bytes(), &got)
			So(err, ShouldBeNil)
			expect.UpdateAt = got.UpdateAt
			soUserShouldEqual(&got, expect)
		})
	})
}
