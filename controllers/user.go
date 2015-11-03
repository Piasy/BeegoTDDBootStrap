package controllers

import (
	"github.com/Piasy/BeegoBootStrap/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Get
// @Description get user by uid
// @Param	phone		path 	string	true		"The phone of the user"
// @Success 200 {object} models.User
// @Failure 403 :phone is empty
// @Failure 404 :phone not found
// @router /:phone [get]
func (this *UserController) Get() {
	phone := this.GetString(":phone")
	if phone != "" {
		beego.Debug("controllers/User::Get() phone: ", phone)
		user, err := models.GetUser(phone)
		beego.Debug("controllers/User::Get() err: ", err)
		if err != nil {
			this.Ctx.ResponseWriter.WriteHeader(404)
			this.Data["json"] = models.ApiError{Code: 1001, Message: err.Error(), Request: this.Ctx.Request.URL.Path}
		} else {
			this.Data["json"] = user
		}
	}
	this.ServeJson()
}

