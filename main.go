package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/Piasy/BeegoBootStrap/docs"
	_ "github.com/Piasy/BeegoBootStrap/routers"
	_ "github.com/Piasy/BeegoBootStrap/models"
)

func init() {
	beego.Debug("main::init() called")

	orm.RegisterDriver("mymysql", orm.DR_MySQL)

	orm.RegisterDataBase("default", "mysql", "root:@/test?charset=utf8")
}

func main() {

	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
