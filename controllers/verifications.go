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
// @Success 201 {object} models.SuccessResult
// @Failure 403 参数错误：缺失或格式错误
// @Failure 500 系统错误
// @router / [post]
func (this *VerificationsController) Post() {
	phone := this.GetString("phone")
	if utils.IsValidPhone(phone) {
		err := models.CreateVerification(phone)
		if err > 0 {
			this.Ctx.ResponseWriter.WriteHeader(500)
			this.Data["json"] = utils.Issue(err, this.Ctx.Request.URL.String())
		} else {
			this.Ctx.ResponseWriter.WriteHeader(201)
			this.Data["json"] = &models.SuccessResult{true}
		}
	} else {
		this.Ctx.ResponseWriter.WriteHeader(403)
		this.Data["json"] = utils.Issue(utils.ERROR_CODE_PARAM_ERROR, this.Ctx.Request.URL.String())
	}
	this.ServeJSON()
}