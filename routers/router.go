// @APIVersion 1.0.0
// @Title ETBoom后端API
package routers

import (
	"github.com/Piasy/BeegoTDDBootStrap/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),
		beego.NSNamespace("/tokens",
			beego.NSInclude(
				&controllers.TokensController{},
			),
		),
		beego.NSNamespace("/verifications",
			beego.NSInclude(
				&controllers.VerificationsController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
