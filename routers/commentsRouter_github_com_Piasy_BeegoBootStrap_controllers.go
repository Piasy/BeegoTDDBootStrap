package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/Piasy/BeegoBootStrap/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/Piasy/BeegoBootStrap/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:phone`,
			[]string{"get"},
			nil})

}
