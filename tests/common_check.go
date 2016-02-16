package test

import (
	"net/http/httptest"
	"bytes"
	"encoding/json"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Piasy/BeegoTDDBootStrap/models"
	"github.com/Piasy/BeegoTDDBootStrap/utils"
)

func soResponseWithStatusCode(w *httptest.ResponseRecorder, code int) {
	So(w.Code, ShouldEqual, code)
	So(w.Body.Len(), ShouldBeGreaterThan, 0)
}

func soEmptyResponseWithStatusCode(w *httptest.ResponseRecorder, code int) {
	So(w.Code, ShouldEqual, code)
	So(w.Body.String(), ShouldEqual, "null")
}

func soUserShouldEqual(actual, expect *models.User) {
	So(actual.Uid, ShouldEqual, expect.Uid)

	So(utils.AreStringEquals(actual.Phone, expect.Phone), ShouldBeTrue)
	So(utils.IsEmptyString(actual.WeiXin), ShouldBeTrue)
	So(utils.IsEmptyString(actual.WeiBo), ShouldBeTrue)
	So(utils.IsEmptyString(actual.QQ), ShouldBeTrue)

	So(actual.Nickname, ShouldEqual, expect.Nickname)
	So(actual.QQNickName, ShouldEqual, expect.QQNickName)
	So(actual.WeiBoNickName, ShouldEqual, expect.WeiBoNickName)
	So(actual.WeiXinNickName, ShouldEqual, expect.WeiXinNickName)
	So(actual.Gender, ShouldEqual, expect.Gender)
	So(actual.Avatar, ShouldEqual, expect.Avatar)

	So(actual.CreateAt, ShouldEqual, expect.CreateAt)
	So(actual.UpdateAt, ShouldEqual, expect.UpdateAt)
}

func soShouldBeApiError(body *bytes.Buffer, code int, request string) {
	var apiError utils.ApiError
	err := json.Unmarshal(body.Bytes(), &apiError)
	So(err, ShouldBeNil)
	So(apiError, ShouldNotBeNil)
	So(apiError.Code, ShouldEqual, code)
	So(apiError.Message, ShouldEqual, utils.ERROR_MESSAGES[code])
	So(apiError.Request, ShouldEqual, request)
}
