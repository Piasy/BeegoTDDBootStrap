package controllers

import (
	"github.com/astaxie/beego"

	"github.com/Piasy/BeegoTDDBootStrap/utils"
	"github.com/Piasy/BeegoTDDBootStrap/models"
)

// 验证码相关接口
type VerificationsController struct {
	beego.Controller
}

// @Title Post
// @Description 请求短信验证码
// @Param	phone		query 	string	true		"手机号"
// @Param	Authorization		header 	string	true		"Basic auth的授权码, 计算方式见wiki"
// @Success 201 {object} models.SuccessResult
// @Failure 401 basic auth失败
// @Failure 403 参数错误：缺失或格式错误
// @Failure 422 手机号已注册
// @Failure 500 系统错误
// @router / [post]
func (this *VerificationsController) Post() {
	phone := this.GetString("phone")
	authorization := this.Ctx.Request.Header.Get("Authorization")
	if authorization != BASIC_AUTH_AUTHORIZATION {
		this.Ctx.ResponseWriter.WriteHeader(401)
		this.Data["json"] = utils.Issue(utils.ERROR_CODE_BASIC_AUTH_FAIL, this.Ctx.Request.URL.String())
	} else if !utils.IsValidPhone(phone) {
		this.Ctx.ResponseWriter.WriteHeader(403)
		this.Data["json"] = utils.Issue(utils.ERROR_CODE_PARAM_ERROR, this.Ctx.Request.URL.String())
	} else if models.UserPhoneExists(&phone) {
		this.Ctx.ResponseWriter.WriteHeader(422)
		this.Data["json"] = utils.Issue(utils.ERROR_CODE_USERS_PHONE_REGISTERED, this.Ctx.Request.URL.String())
	} else if err := models.CreateVerification(phone); err > 0 {
		this.Ctx.ResponseWriter.WriteHeader(500)
		this.Data["json"] = utils.Issue(err, this.Ctx.Request.URL.String())
	} else {
		this.Ctx.ResponseWriter.WriteHeader(201)
		this.Data["json"] = &models.SuccessResult{true}
	}
	this.ServeJSON()
}
