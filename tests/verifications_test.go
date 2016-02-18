package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/stretchr/testify/assert"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/Piasy/BeegoTDDBootStrap/utils"
	"github.com/Piasy/BeegoTDDBootStrap/models"
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

func createVerification(t *testing.T, subject, phone string) *models.Verification {
	request, _ := http.NewRequest("POST", "/v1/verifications/", bytes.NewBuffer([]byte("phone=" + phone)))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", "dGVzdF9jbGllbnQ6dGVzdF9wYXNz")
	recorder := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(recorder, request)
	beego.Debug("testing <", subject, ">: create verification, Code[", recorder.Code, "]\n", recorder.Body.String())

	Convey("Subject: " + subject + ": create verification\n", t, func() {
		Convey("Status code should be 201", func() {
			soEmptyResponseWithStatusCode(recorder, 201)
		})
	})

	o := orm.NewOrm()
	verification := models.Verification{Phone: phone}
	err := o.Read(&verification, "Phone")
	assert.Nil(t, err)

	return &verification
}

func deleteVerification(t *testing.T, id int64) {
	o := orm.NewOrm()
	verification := models.Verification{Id: id}
	_, err := o.Delete(&verification)
	assert.Nil(t, err)
}
