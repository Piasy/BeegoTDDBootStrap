package models_test

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

var ormInitiated bool = false

func initORM() {
	// switch to prod
	beego.BConfig.RunMode = "prod"
	if ormInitiated {
		return
	}
	appConf, err := config.NewConfig("ini", "../conf/app.conf")
	if err != nil {
		panic(err)
	}
	dbAddr := appConf.String("admin::dbAddr")
	dbUser := appConf.String("admin::dbUser")
	dbPass := appConf.String("admin::dbPass")

	orm.RegisterDriver("mymysql", orm.DRMySQL)
	conn := fmt.Sprintf("%s:%s@tcp(%s)/beego_unit_test?charset=utf8mb4", dbUser, dbPass, dbAddr)
	orm.RegisterDataBase("default", "mysql", conn)
	ormInitiated = true
}
