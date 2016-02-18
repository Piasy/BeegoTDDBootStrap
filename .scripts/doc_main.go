package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/Piasy/HabitsAPI/docs"
	_ "github.com/Piasy/HabitsAPI/routers"
	_ "github.com/Piasy/HabitsAPI/models"
	"github.com/Piasy/HabitsAPI/controllers"
	"github.com/Piasy/HabitsAPI/utils"
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
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.Run()
}
