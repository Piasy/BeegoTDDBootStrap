package controllers

import (
	"github.com/astaxie/beego"

	"github.com/Piasy/BeegoTDDBootStrap/models"
	"github.com/Piasy/BeegoTDDBootStrap/utils"
)

// Token相关接口
type TokensController struct {
	beego.Controller
}

// @Title Post
// @Description 通过手机号和密码获取Token（登录）
// @Param	phone		query 	string	true		"用户手机号"
// @Param	secret		query 	string	true		"加密处理后的密码，全部小写"
// @Success 201 {object} models.User
// @Failure 403 参数错误：缺失或格式错误
// @Failure 422 手机号未注册/密码错误
// @Failure 500 系统错误
// @router / [post]
func (this *TokensController) Post() {
	phone := this.GetString("phone")
	secret := this.GetString("secret")
	if utils.IsValidPhone(phone) && len(secret) == 40 {
		user, err := models.VerifyUserByPhone(phone, secret)
		if err > 0 {
			this.Ctx.ResponseWriter.WriteHeader(422)
			this.Data["json"] = utils.Issue(err, this.Ctx.Request.URL.String())
		} else {
			this.Ctx.ResponseWriter.WriteHeader(201)
			this.Data["json"] = user
		}
	} else {
		this.Ctx.ResponseWriter.WriteHeader(403)
		this.Data["json"] = utils.Issue(utils.ERROR_CODE_PARAM_ERROR, this.Ctx.Request.URL.String())
	}
	this.ServeJson()
}