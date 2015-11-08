package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/Piasy/BeegoTDDBootStrap/controllers:TokensController"] = append(beego.GlobalControllerRouter["github.com/Piasy/BeegoTDDBootStrap/controllers:TokensController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/Piasy/BeegoTDDBootStrap/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/Piasy/BeegoTDDBootStrap/controllers:UsersController"],
		beego.ControllerComments{
			"Get",
			`/?:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/Piasy/BeegoTDDBootStrap/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/Piasy/BeegoTDDBootStrap/controllers:UsersController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/Piasy/BeegoTDDBootStrap/controllers:VerificationsController"] = append(beego.GlobalControllerRouter["github.com/Piasy/BeegoTDDBootStrap/controllers:VerificationsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

}
