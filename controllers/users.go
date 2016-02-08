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

// @Title Post
// @Description 通过手机号注册, 返回所有字段
// @Param	phone		query 	string	true		"用户手机号"
// @Param	code		query 	string	true		"手机验证码"
// @Param	secret		query 	string	true		"加密处理后的密码"
// @Param	Authorization		header 	string	true		"Basic auth的授权码, 计算方式见wiki"
// @Success 201 {object} models.User
// @Failure 401 basic auth失败
// @Failure 403 参数错误：缺失或格式错误
// @Failure 422 手机号已注册
// @Failure 500 系统错误
// @router / [post]
func (this *UsersController) Post() {
	phone := this.GetString("phone")
	code := this.GetString("code")
	secret := this.GetString("secret")
	authorization := this.Ctx.Request.Header.Get("Authorization")
	if authorization != BASIC_AUTH_AUTHORIZATION {
		this.Ctx.ResponseWriter.WriteHeader(401)
		this.Data["json"] = utils.Issue(utils.ERROR_CODE_BASIC_AUTH_FAIL, this.Ctx.Request.URL.String())
	} else if !utils.IsValidPhone(phone) || code == "" || len(secret) != 40 {
		this.Ctx.ResponseWriter.WriteHeader(403)
		this.Data["json"] = utils.Issue(utils.ERROR_CODE_PARAM_ERROR, this.Ctx.Request.URL.String())
	} else if exists := models.UserPhoneExists(&phone); exists {
		this.Ctx.ResponseWriter.WriteHeader(422)
		this.Data["json"] = utils.Issue(utils.ERROR_CODE_USERS_PHONE_REGISTERED, this.Ctx.Request.URL.String())
	} else if err := models.CheckVerifyCode(phone, code); err > 0 {
		this.Ctx.ResponseWriter.WriteHeader(422)
		this.Data["json"] = utils.Issue(err, this.Ctx.Request.URL.String())
	} else if user, err := models.CreateUserByPhone(&phone, secret); err > 0 {
		this.Ctx.ResponseWriter.WriteHeader(500)
		this.Data["json"] = utils.Issue(err, this.Ctx.Request.URL.String())
	} else {
		this.Ctx.ResponseWriter.WriteHeader(201)
		this.Data["json"] = user
	}

	this.ServeJSON()
}

// @Title Get
// @Description [允许游客]通过uid获取用户信息; <br>请求自己的信息返回所有字段; <br>请求他人信息只有uid, nickname, avatar, gender字段; <br>其中gender 0表示未知, 1表示男, 2表示女
// @Param	uid		path 	int64	true		"目标用户uid"
// @Param	token		query 	string	true		"自己的token"
// @Success 200 {object} models.User
// @Failure 401 token无效
// @Failure 403 参数错误：非法uid
// @Failure 404 目标用户不存在
// @router /:uid [get]
func (this *UsersController) Get() {
	uid, err := this.GetInt64(":uid")
	token := this.GetString("token")
	if token == "" {
		this.Ctx.ResponseWriter.WriteHeader(401)
		this.Data["json"] = utils.Issue(utils.ERROR_CODE_TOKENS_INVALID_TOKEN, this.Ctx.Request.URL.String())
	} else if err != nil || uid < utils.USER_MIN_UID {
		this.Ctx.ResponseWriter.WriteHeader(403)
		this.Data["json"] = utils.Issue(utils.ERROR_CODE_PARAM_ERROR, this.Ctx.Request.URL.String())
	} else if token == VISITOR_TOKEN {
		if user, errNum := models.GetUserByUid(uid); errNum > 0 {
			this.Ctx.ResponseWriter.WriteHeader(404)
			this.Data["json"] = utils.Issue(errNum, this.Ctx.Request.URL.String())
		} else {
			this.Ctx.ResponseWriter.WriteHeader(200)
			ret := models.User{Uid: user.Uid, Nickname: user.Nickname, Avatar: user.Avatar,
				Gender: user.Gender}
			this.Data["json"] = &ret
		}
	} else if self, errNum := models.GetUserByToken(token); errNum > 0 {
		this.Ctx.ResponseWriter.WriteHeader(401)
		this.Data["json"] = utils.Issue(errNum, this.Ctx.Request.URL.String())
	} else if self.Uid == uid {
		this.Ctx.ResponseWriter.WriteHeader(200)
		this.Data["json"] = self
	} else if user, errNum := models.GetUserByUid(uid); errNum > 0 {
		this.Ctx.ResponseWriter.WriteHeader(404)
		this.Data["json"] = utils.Issue(errNum, this.Ctx.Request.URL.String())
	} else {
		this.Ctx.ResponseWriter.WriteHeader(200)
		ret := models.User{Uid: user.Uid, Nickname: user.Nickname, Avatar: user.Avatar,
			Gender: user.Gender}
		this.Data["json"] = &ret
	}
	this.ServeJSON()
}

// @Title Patch
// @Description 修改用户自己的信息, 修改哪些字段就传哪些字段, 成功后返回所有字段, <br/>注意: N位中英数限制, 指只能有汉字/英文字母/阿拉伯数字, 不能有标点符号, 特殊符号, 每个汉字/字母/数字长度都算1
// @Param	token		query 	string	true		"Token"
// @Param	phone		query 	string	false		"更换绑定手机号"
// @Param	code		query 	string	false		"手机验证码, 换绑手机号时需要"
// @Param	wx_openid		query 	string	false		"微信授权的openid"
// @Param	wx_token		query 	string	false		"微信授权的token"
// @Param	qq_openid		query 	string	false		"QQ授权的openid"
// @Param	qq_token		query 	string	false		"QQ授权的token"
// @Param	wb_token		query 	string	false		"微博授权的token"
// @Param	nickname		query 	string	false		"昵称, 小于12位中英数"
// @Param	gender		query 	int	false		"性别, 1为男, 2为女"
// @Param	avatar		query 	string	false		"头像url, 通过上传头像接口上传成功后获得"
// @Success 201 {object} models.User
// @Failure 401 token无效
// @Failure 403 参数错误：缺失或格式错误
// @Failure 500 系统错误
// @router / [patch]
func (this *UsersController) Patch() {
	token := this.GetString("token")
	phone := this.GetString("phone")
	code := this.GetString("code")
	wx_openid := this.GetString("wx_openid")
	wx_token := this.GetString("wx_token")
	qq_openid := this.GetString("qq_openid")
	qq_token := this.GetString("qq_token")
	wb_token := this.GetString("wb_token")
	nickname := this.GetString("nickname")
	gender, errGender := this.GetInt("gender", 0)
	avatar := this.GetString("avatar")
	if (phone != "" && !utils.IsValidPhone(phone)) || errGender != nil {
		// has phone, but invalid; parse gender/birthday error
		this.Ctx.ResponseWriter.WriteHeader(403)
		this.Data["json"] = utils.Issue(utils.ERROR_CODE_PARAM_ERROR, this.Ctx.Request.URL.String())
	} else if user, err := models.GetUserByToken(token); err > 0 {
		// invalid token
		this.Ctx.ResponseWriter.WriteHeader(401)
		this.Data["json"] = utils.Issue(err, this.Ctx.Request.URL.String())
	} else {
		for {
			// has valid phone
			if phone != "" {
				if code == "" {
					// has valid phone but not valid code
					this.Ctx.ResponseWriter.WriteHeader(403)
					this.Data["json"] = utils.Issue(utils.ERROR_CODE_VERIFY_CODE_MISMATCH, this.Ctx.Request.URL.String())
					break

				}
				if err = models.CheckVerifyCode(phone, code); err > 0 {
					// code mismatch
					this.Ctx.ResponseWriter.WriteHeader(403)
					this.Data["json"] = utils.Issue(err, this.Ctx.Request.URL.String())
					break
				}
				// valid
				user.Phone = &phone
			}
			// has wx_openid
			if wx_openid != "" {
				if wx_token == "" {
					// empty wx_token
					this.Ctx.ResponseWriter.WriteHeader(403)
					this.Data["json"] = utils.Issue(utils.ERROR_CODE_PARAM_ERROR, this.Ctx.Request.URL.String())
					break
				}
				authUser, err := utils.AuthWithWeiXin(wx_openid, wx_token)
				if err > 0 {
					// auth fail
					this.Ctx.ResponseWriter.WriteHeader(403)
					this.Data["json"] = utils.Issue(err, this.Ctx.Request.URL.String())
					break
				}
				// verified
				user.WeiXin = &authUser.Openid
				user.WeiXinNickName = authUser.Nickname
			}
			// has wb_token
			if wb_token != "" {
				authUser, err := utils.AuthWithWeiBo(wb_token)
				if err > 0 {
					// auth fail
					this.Ctx.ResponseWriter.WriteHeader(403)
					this.Data["json"] = utils.Issue(err, this.Ctx.Request.URL.String())
					break
				}
				// verified
				user.WeiBo = &authUser.Openid
				user.WeiBoNickName = authUser.Nickname
			}
			// has qq_openid
			if qq_openid != "" {
				if qq_token == "" {
					// empty qq_token
					this.Ctx.ResponseWriter.WriteHeader(403)
					this.Data["json"] = utils.Issue(utils.ERROR_CODE_PARAM_ERROR, this.Ctx.Request.URL.String())
					break
				}
				authUser, err := utils.AuthWithQQ(qq_openid, qq_token, QQ_OAUTH_CONSUMER_KEY)
				if err > 0 {
					// auth fail
					this.Ctx.ResponseWriter.WriteHeader(403)
					this.Data["json"] = utils.Issue(err, this.Ctx.Request.URL.String())
					break
				}
				// verified
				user.QQ = &authUser.Openid
				user.QQNickName = authUser.Nickname
			}
			if nickname != "" {
				if !utils.IsLegalRestrictedStringWithLength(nickname, utils.USER_NICKNAME_MEX_LEN) {
					this.Ctx.ResponseWriter.WriteHeader(403)
					this.Data["json"] = utils.Issue(utils.ERROR_CODE_USERS_INVALID_NICKNAME, this.Ctx.Request.URL.String())
					break
				}
				user.Nickname = nickname
			}
			if gender > 0 {
				if gender != 1 && gender != 2 {
					this.Ctx.ResponseWriter.WriteHeader(403)
					this.Data["json"] = utils.Issue(utils.ERROR_CODE_USERS_INVALID_GENDER_VALUE, this.Ctx.Request.URL.String())
					break
				}
				user.Gender = gender
			}
			if avatar != "" {
				if len(avatar) > utils.USER_AVATAR_MEX_LEN {
					this.Ctx.ResponseWriter.WriteHeader(403)
					this.Data["json"] = utils.Issue(utils.ERROR_CODE_USERS_INVALID_AVATAR, this.Ctx.Request.URL.String())
					break
				}
				user.Avatar = avatar
			}

			err = models.UpdateUser(user)
			if err > 0 {
				this.Ctx.ResponseWriter.WriteHeader(403)
				this.Data["json"] = utils.Issue(err, this.Ctx.Request.URL.String())
				break
			}

			// success
			this.Ctx.ResponseWriter.WriteHeader(201)
			this.Data["json"] = user
			break
		}
	}
	this.ServeJSON()
}
