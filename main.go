package main

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/Piasy/BeegoTDDBootStrap/docs"
	_ "github.com/Piasy/BeegoTDDBootStrap/routers"
	_ "github.com/Piasy/BeegoTDDBootStrap/models"
	"github.com/Piasy/BeegoTDDBootStrap/controllers"
	"github.com/Piasy/BeegoTDDBootStrap/utils"
)

func init() {
	appConf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	dbAddr := appConf.String("admin::dbAddr")
	dbUser := appConf.String("admin::dbUser")
	dbPass := appConf.String("admin::dbPass")
	dbName := appConf.String("admin::dbName")
	controllers.ALI_YUN_AK_ID = appConf.String("admin::akId")
	controllers.ALI_YUN_AK_KEY = appConf.String("admin::akKey")
	controllers.QQ_OAUTH_CONSUMER_KEY = appConf.String("admin::qqOAuthConsumerKey")
	clientId := appConf.String("admin::clientId")
	clientSecret := appConf.String("admin::clientSecret")
	controllers.BASIC_AUTH_AUTHORIZATION = utils.Base64(clientId + ":" + clientSecret)

	controllers.VISITOR_TOKEN = appConf.String("admin::visitorToken")

	orm.RegisterDriver("mymysql", orm.DRMySQL)
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", dbUser, dbPass, dbAddr, dbName)
	orm.RegisterDataBase("default", "mysql", conn)
}

func main() {
	beego.SetLogger("file", `{"filename":"logs/beego_bootstrap_api.log"}`)
	orm.RunCommand()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

		orm.Debug = true
		file, _ := os.OpenFile("logs/orm.log", os.O_APPEND|os.O_WRONLY, 0600)
		orm.DebugLog = orm.NewLog(file)
	}
	beego.Run()
}
