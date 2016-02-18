package test

import (
	"net/http"
	"net/http/httptest"
	"runtime"
	"path/filepath"
	"fmt"
	"testing"
	"strconv"

	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "github.com/smartystreets/goconvey/convey"

	_ "github.com/Piasy/BeegoTDDBootStrap/routers"
	"github.com/Piasy/BeegoTDDBootStrap/controllers"
	"github.com/Piasy/BeegoTDDBootStrap/utils"
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
	dbAddr := appConf.String("admin::dbAddr")
	dbUser := appConf.String("admin::dbUser")
	dbPass := appConf.String("admin::dbPass")
	controllers.ALI_YUN_AK_ID = appConf.String("admin::akId")
	controllers.ALI_YUN_AK_KEY = appConf.String("admin::akKey")
	controllers.QQ_OAUTH_CONSUMER_KEY = appConf.String("admin::qqOAuthConsumerKey")
	clientId := appConf.String("admin::clientId")
	clientSecret := appConf.String("admin::clientSecret")
	controllers.BASIC_AUTH_AUTHORIZATION = utils.Base64(clientId + ":" + clientSecret)

	controllers.VISITOR_TOKEN = appConf.String("admin::visitorToken")

	orm.RegisterDriver("mymysql", orm.DRMySQL)
	conn := fmt.Sprintf("%s:%s@tcp(%s)/beego_unit_test?charset=utf8mb4", dbUser, dbPass, dbAddr)
	orm.RegisterDataBase("default", "mysql", conn)
	ormInitiated = true
}

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	initORM()
	// switch to prod
	beego.BConfig.RunMode = "prod"
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
