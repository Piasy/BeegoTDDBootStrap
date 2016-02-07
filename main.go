package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/Piasy/BeegoTDDBootStrap/docs"
	_ "github.com/Piasy/BeegoTDDBootStrap/routers"
	_ "github.com/Piasy/BeegoTDDBootStrap/models"
)

func init() {
	appConf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	dbUser := appConf.String("admin::dbUser")
	dbPass := appConf.String("admin::dbPass")
	dbName := appConf.String("admin::dbName")

	orm.RegisterDriver("mymysql", orm.DRMySQL)

	var conn string
	if dbPass == "" {
		conn = fmt.Sprintf("%s:@/%s?charset=utf8", dbUser, dbName)
	} else {
		conn = fmt.Sprintf("%s:%s@/%s?charset=utf8", dbUser, dbPass, dbName)
	}
	orm.RegisterDataBase("default", "mysql", conn)
}

func main() {
	beego.SetLogger("file", `{"filename":"logs/etboom.log"}`)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
