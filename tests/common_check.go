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

func soUserShouldEqual(actual, expect *models.User) {
	So(actual.Uid, ShouldEqual, expect.Uid)
	So(actual.Username, ShouldEqual, expect.Username)
	So(actual.Phone, ShouldEqual, expect.Phone)
	So(actual.Nickname, ShouldEqual, expect.Nickname)
	So(actual.Phone, ShouldEqual, expect.Phone)
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
