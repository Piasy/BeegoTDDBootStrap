package controllers

import (
	"github.com/astaxie/beego"

	"github.com/Piasy/BeegoTDDBootStrap/models"
	"github.com/Piasy/BeegoTDDBootStrap/utils"
)

// 用户系统相关接口
type UsersController struct {
	beego.Controller
}

// @Title Get
// @Description 通过uid获取用户信息, 请求自己的信息返回所有字段, 请求他人信息只有uid和nickname字段
// @Param	uid		path 	int64	true		"目标用户uid"
// @Param	token		query 	string	true		"自己的token"
// @Success 200 {object} models.User
// @Failure 403 参数错误：非法uid，token错误
// @Failure 404 目标用户不存在
// @router /?:uid [get]
func (this *UsersController) Get() {
	uid, err := this.GetInt64(":uid")
	token := this.GetString("token")
	if err == nil && uid > utils.USER_MIN_UID && token != "" {
		self, errNum := models.GetUserByToken(token)
		if errNum > 0 {
			this.Ctx.ResponseWriter.WriteHeader(403)
			this.Data["json"] = utils.Issue(errNum, this.Ctx.Request.URL.String())
		} else {
			if self.Uid == uid {
				this.Ctx.ResponseWriter.WriteHeader(200)
				this.Data["json"] = self
			} else {
				user, errNum := models.GetUserByUid(uid)
				if errNum > 0 {
					this.Ctx.ResponseWriter.WriteHeader(404)
					this.Data["json"] = utils.Issue(errNum, this.Ctx.Request.URL.String())
				} else {
					this.Ctx.ResponseWriter.WriteHeader(200)
					ret := models.User{Uid: user.Uid, Nickname: user.Nickname}
					this.Data["json"] = &ret
				}
			}
		}
	} else {
		this.Ctx.ResponseWriter.WriteHeader(403)
		this.Data["json"] = utils.Issue(utils.ERROR_CODE_PARAM_ERROR, this.Ctx.Request.URL.String())
	}
	this.ServeJSON()
}

// @Title Post
// @Description 通过手机号注册, 返回所有字段
// @Param	phone		query 	string	true		"用户手机号"
// @Param	code		query 	string	true		"手机验证码"
// @Param	secret		query 	string	true		"加密处理后的密码"
// @Success 201 {object} models.User
// @Failure 403 参数错误：缺失或格式错误
// @Failure 422 手机号已注册
// @Failure 500 系统错误
// @router / [post]
func (this *UsersController) Post() {
	phone := this.GetString("phone")
	code := this.GetString("code")
	secret := this.GetString("secret")

	if utils.IsValidPhone(phone) && code != "" && len(secret) == 40 {
		exists := models.UserPhoneExists(phone)
		if exists {
			this.Ctx.ResponseWriter.WriteHeader(422)
			this.Data["json"] = utils.Issue(utils.ERROR_CODE_USERS_PHONE_REGISTERED, this.Ctx.Request.URL.String())
		} else if err := models.CheckVerifyCode(phone, code); err == 0 {
			user, err := models.CreateUserByPhone(phone, secret)
			if err > 0 {
				this.Ctx.ResponseWriter.WriteHeader(500)
				this.Data["json"] = utils.Issue(err, this.Ctx.Request.URL.String())
			} else {
				this.Ctx.ResponseWriter.WriteHeader(201)
				this.Data["json"] = user
			}
		} else {
			this.Ctx.ResponseWriter.WriteHeader(422)
			this.Data["json"] = utils.Issue(err, this.Ctx.Request.URL.String())
		}
	} else {
		this.Ctx.ResponseWriter.WriteHeader(403)
		this.Data["json"] = utils.Issue(utils.ERROR_CODE_PARAM_ERROR, this.Ctx.Request.URL.String())
	}
	this.ServeJSON()
}
