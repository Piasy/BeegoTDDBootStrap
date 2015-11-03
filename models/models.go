package models

import (
	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

func init() {
	beego.Debug("models/models::init() called")
	orm.RegisterModel(new(User))
}
