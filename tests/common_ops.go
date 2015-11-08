package test

import (
	"net/http"
	"net/http/httptest"
	"runtime"
	"path/filepath"
	"fmt"
	"testing"
	"encoding/json"
	"bytes"
	"strconv"

	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/stretchr/testify/assert"
	. "github.com/smartystreets/goconvey/convey"

	_ "github.com/Piasy/BeegoTDDBootStrap/routers"
	"github.com/Piasy/BeegoTDDBootStrap/models"
)

var ormInitiated bool = false

func initORM() {
	if ormInitiated {
		return
	}
	appConf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	dbUser := appConf.String("admin::dbUser")
	dbPass := appConf.String("admin::dbPass")
	dbName := "beego_unit_test"

	orm.RegisterDriver("mymysql", orm.DR_MySQL)

	var conn string
	if dbPass == "" {
		conn = fmt.Sprintf("%s:@/%s?charset=utf8", dbUser, dbName)
	} else {
		conn = fmt.Sprintf("%s:%s@/%s?charset=utf8", dbUser, dbPass, dbName)
	}
	orm.RegisterDataBase("default", "mysql", conn)
	ormInitiated = true
}

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	initORM()
	// switch to prod
	beego.RunMode = "prod"
}

func deleteUser(t *testing.T, id int64) {
	o := orm.NewOrm()
	user := models.User{Id: id}
	_, err := o.Delete(&user)
	assert.Nil(t, err)
}

func deleteVerification(t *testing.T, id int64) {
	o := orm.NewOrm()
	verification := models.Verification{Id: id}
	_, err := o.Delete(&verification)
	assert.Nil(t, err)
}

func createVerification(t *testing.T, subject, phone string) *models.Verification  {
	request, _ := http.NewRequest("POST", "/v1/verifications/", bytes.NewBuffer([]byte("phone=" + phone)))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(recorder, request)
	beego.Debug("testing <", subject, ">: create verification, Code[", recorder.Code, "]\n", recorder.Body.String())

	Convey("Subject: " + subject + ": create verification\n", t, func() {
		Convey("Status code should be 201", func() {
			soResponseWithStatusCode(recorder, 201)
		})
		Convey("Create should success", func() {
			var success models.SuccessResult
			err := json.Unmarshal(recorder.Body.Bytes(), &success)
			So(err, ShouldBeNil)
			So(success.Success, ShouldBeTrue)
		})
	})

	o := orm.NewOrm()
	verification := models.Verification{Phone: phone}
	err := o.Read(&verification, "Phone")
	assert.Nil(t, err)

	return &verification
}

func createUser(t *testing.T, subject, phone, secret, code string) *models.User {
	request, _ := http.NewRequest("POST", "/v1/users/",
		bytes.NewBuffer([]byte(fmt.Sprintf("phone=%s&secret=%s&code=%s", phone, secret, code))))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(recorder, request)
	beego.Debug("testing <", subject, ">: create user, Code[", recorder.Code, "]\n", recorder.Body.String())

	o := orm.NewOrm()
	fromDB := models.User{Phone: "18801234567"}
	err := o.Read(&fromDB, "Phone")
	assert.Nil(t, err)

	Convey("Subject: " + subject + ": create user\n", t, func() {
		Convey("Status code should be 201", func() {
			soResponseWithStatusCode(recorder, 201)
		})
		Convey("Create should success", func() {
			var created models.User
			err := json.Unmarshal(recorder.Body.Bytes(), &created)
			So(err, ShouldBeNil)
			soUserShouldEqual(&created, &fromDB)
		})
	})

	return &fromDB
}

func checkHasUserByUid(t *testing.T, subject string, expect *models.User) {
	request, _ := http.NewRequest("GET", fmt.Sprintf("/v1/users/%d?token=%s", expect.Uid, expect.Token), nil)
	recorder := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(recorder, request)
	beego.Debug("testing <", subject, ">: get user, Code[", recorder.Code, "]\n", recorder.Body.String())

	Convey("Subject: " + subject + ": get user\n", t, func() {
		Convey("Get user status code should be 200", func() {
			soResponseWithStatusCode(recorder, 200)
		})
		Convey("Get user should be same as created", func() {
			var got models.User
			err := json.Unmarshal(recorder.Body.Bytes(), &got)
			So(err, ShouldBeNil)
			soUserShouldEqual(&got, expect)
		})
	})
}

func checkIsApiErrorRequest(t *testing.T, subject string, request *http.Request, statusCode, errorCode int) {
	recorder := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(recorder, request)
	beego.Debug("testing <", subject, ">, Code[", recorder.Code, "]\n", recorder.Body.String())

	Convey("Subject: " + subject + "\n", t, func() {
		Convey("Status code should be " + strconv.Itoa(statusCode), func() {
			soResponseWithStatusCode(recorder, statusCode)
		})
		Convey("The result should be an ApiError", func() {
			soShouldBeApiError(recorder.Body, errorCode, request.URL.String())
		})
	})
}

func updateTokenByPhone(t *testing.T, subject string, expect *models.User, secret string) {
	request, _ := http.NewRequest("POST", "/v1/tokens/",
		bytes.NewBuffer([]byte(fmt.Sprintf("phone=%s&secret=%s", expect.Phone, secret))))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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
			soUserShouldEqual(&got, expect)
		})
	})
}