package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/Piasy/BeegoTDDBootStrap/utils"
)

type User struct {
	Id             int64 `json:"-" orm:"pk;auto"`

	Uid            int64 `json:"uid" orm:"column(uid);unique;index"`
	Token          string `json:"token,omitempty" orm:"column(token);unique;index;size(40)"`
	Phone          *string `json:"phone,omitempty" orm:"column(phone);null;unique;index;size(11)"`
	WeiXin         *string `json:"-" orm:"column(weixin);null;unique;index;size(191)"`
	WeiBo          *string `json:"-" orm:"column(weibo);null;unique;index;size(191)"`
	QQ             *string `json:"-" orm:"column(qq);null;unique;index;size(191)"`

	Password       string `json:"-" orm:"column(password);size(40)"`

	Nickname       string `json:"nickname,omitempty" orm:"column(nickname);size(12)"`
	QQNickName     string `json:"qq_nickname,omitempty" orm:"column(qq_nickname);size(127)"`
	WeiBoNickName  string `json:"weibo_nickname,omitempty" orm:"column(weibo_nickname);size(127)"`
	WeiXinNickName string `json:"weixin_nickname,omitempty" orm:"column(weixin_nickname);size(127)"`

	Gender         int `json:"gender" orm:"column(gender)"`
	Avatar         string `json:"avatar,omitempty" orm:"column(avatar);size(191)"`

	CreateAt       int64 `json:"create_at" orm:"column(create_at);unique"`
	UpdateAt       int64 `json:"update_at" orm:"column(update_at);unique"`
}

const USERS_TABLE_NAME string = "users"

func (u *User) TableName() string {
	return USERS_TABLE_NAME
}

func UserPhoneExists(phone *string) bool {
	utils.AssertNotEmptyString(phone)
	o := orm.NewOrm()
	user := User{Phone: phone}
	err := o.Read(&user, "Phone")
	return err == nil
}

func GetUserByUid(uid int64) (*User, int) {
	o := orm.NewOrm()
	return getUserByUidInternal(&o, uid)
}

func getUserByUidInternal(o *orm.Ormer, uid int64) (*User, int) {
	user := User{Uid: uid}
	err := (*o).Read(&user, "Uid")
	if err == nil {
		return &user, 0
	}
	return nil, utils.ERROR_CODE_USERS_USER_NOT_EXISTS
}

func GetUserByToken(token string) (*User, int) {
	o := orm.NewOrm()
	user := User{Token: token}
	err := o.Read(&user, "Token")
	if err == nil {
		return &user, 0
	}
	return nil, utils.ERROR_CODE_TOKENS_INVALID_TOKEN
}

func AuthWithWeiXin(openid, token string) (*User, int) {
	authUser, errNum := utils.AuthWithWeiXin(openid, token)
	if errNum > 0 {
		return nil, errNum
	}
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		beego.Warning("AuthWithWeiXin fail: ", err)
		return nil, utils.ERROR_CODE_SYSTEM_ERROR
	}

	user := getUserByWeiXinInternal(&o, &authUser.Openid)
	user, errNum = createOrUpdateUserInternal(&o, user, authUser, "AuthWithWeiXin fail: ",
		utils.SNS_PLATFORM_WEIXIN)
	if errNum > 0 {
		o.Rollback()
		return nil, errNum
	}
	err = o.Commit()
	if err != nil {
		beego.Warning("AuthWithWeiXin fail: commit fail ", err)
		o.Rollback()
		return nil, utils.ERROR_CODE_SYSTEM_ERROR
	}

	return user, 0
}

func getUserByWeiXinInternal(o *orm.Ormer, openid *string) *User {
	utils.AssertNotEmptyString(openid)
	user := User{WeiXin: openid}
	err := (*o).Read(&user, "WeiXin")
	if err == nil {
		return &user
	}
	return nil
}

func AuthWithWeiBo(token string) (*User, int) {
	authUser, errNum := utils.AuthWithWeiBo(token)
	if errNum > 0 {
		return nil, errNum
	}
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		beego.Warning("AuthWithWeiBo fail: ", err)
		return nil, utils.ERROR_CODE_SYSTEM_ERROR
	}

	user := getUserByWeiBoInternal(&o, &authUser.Openid)
	user, errNum = createOrUpdateUserInternal(&o, user, authUser, "AuthWithWeiBo fail: ",
		utils.SNS_PLATFORM_WEIBO)
	if errNum > 0 {
		o.Rollback()
		return nil, errNum
	}
	err = o.Commit()
	if err != nil {
		beego.Warning("AuthWithWeiBo fail: commit fail ", err)
		o.Rollback()
		return nil, utils.ERROR_CODE_SYSTEM_ERROR
	}

	return user, 0
}

func getUserByWeiBoInternal(o *orm.Ormer, openid *string) *User {
	utils.AssertNotEmptyString(openid)
	user := User{WeiBo: openid}
	err := (*o).Read(&user, "WeiBo")
	if err == nil {
		return &user
	}
	return nil
}

func AuthWithQQ(openid, token, appId string) (*User, int) {
	authUser, errNum := utils.AuthWithQQ(openid, token, appId)
	if errNum > 0 {
		return nil, errNum
	}
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		beego.Warning("AuthWithQQ fail: ", err)
		return nil, utils.ERROR_CODE_SYSTEM_ERROR
	}

	user := getUserByQQInternal(&o, &authUser.Openid)
	user, errNum = createOrUpdateUserInternal(&o, user, authUser, "AuthWithQQ fail: ",
		utils.SNS_PLATFORM_QQ)
	if errNum > 0 {
		o.Rollback()
		return nil, errNum
	}
	err = o.Commit()
	if err != nil {
		beego.Warning("AuthWithQQ fail: commit fail ", err)
		o.Rollback()
		return nil, utils.ERROR_CODE_SYSTEM_ERROR
	}

	return user, 0
}

func getUserByQQInternal(o *orm.Ormer, openid *string) *User {
	utils.AssertNotEmptyString(openid)
	user := User{QQ: openid}
	err := (*o).Read(&user, "QQ")
	if err == nil {
		return &user
	}
	return nil
}

// callee's duty to commit & rollback
func createOrUpdateUserInternal(o *orm.Ormer, user *User, authUser *utils.AuthUserInfo, logTag string, platform int) (*User, int) {
	if user != nil {
		if utils.IsLegalRestrictedStringWithLength(authUser.Nickname, utils.USER_NICKNAME_MEX_LEN) {
			user.Nickname = authUser.Nickname
		}
		user.Avatar = authUser.Avatar
		user.Gender = authUser.Gender
		switch platform {
		case utils.SNS_PLATFORM_WEIXIN:
			user.WeiXinNickName = authUser.Nickname
		case utils.SNS_PLATFORM_WEIBO:
			user.WeiBoNickName = authUser.Nickname
		case utils.SNS_PLATFORM_QQ:
			user.QQNickName = authUser.Nickname
		}

		var err int
		for i := 0; i < DB_UNIQUE_CONFLICT_TRY; i++ {
			user.Token = utils.GenToken()
			user.UpdateAt = utils.GetTimeMillis()
			err = updateUserInternal(o, user)
			if err == 0 {
				return user, 0
			}
			time.Sleep(1 * time.Millisecond)
		}

		beego.Warning(logTag, err)
		return nil, utils.ERROR_CODE_SYSTEM_ERROR
	}

	user = &User{Gender: authUser.Gender, Avatar: authUser.Avatar}
	if utils.IsLegalRestrictedStringWithLength(authUser.Nickname, utils.USER_NICKNAME_MEX_LEN) {
		user.Nickname = authUser.Nickname
	}
	switch platform {
	case utils.SNS_PLATFORM_WEIXIN:
		user.WeiXin = &authUser.Openid
		user.WeiXinNickName = authUser.Nickname
	case utils.SNS_PLATFORM_WEIBO:
		user.WeiBo = &authUser.Openid
		user.WeiBoNickName = authUser.Nickname
	case utils.SNS_PLATFORM_QQ:
		user.QQ = &authUser.Openid
		user.QQNickName = authUser.Nickname
	}

	var err error
	for i := 0; i < DB_UNIQUE_CONFLICT_TRY; i++ {
		user.Uid = utils.GenUid()
		user.Token = utils.GenToken()
		now := utils.GetTimeMillis()
		user.CreateAt = now
		user.UpdateAt = now
		_, err = (*o).Insert(user)
		if err == nil {
			return user, 0
		}
		time.Sleep(1 * time.Millisecond)
	}

	beego.Warning(logTag, err)
	return nil, utils.ERROR_CODE_SYSTEM_ERROR
}

func CreateUserByPhone(phone *string, secret string) (*User, int) {
	utils.AssertNotEmptyString(phone)
	o := orm.NewOrm()
	password := utils.Secret2Password("phone:" + *phone, secret)
	user := User{Phone: phone, Password: password}

	var err error
	for i := 0; i < DB_UNIQUE_CONFLICT_TRY; i++ {
		user.Uid = utils.GenUid()
		user.Token = utils.GenToken()
		now := utils.GetTimeMillis()
		user.CreateAt = now
		user.UpdateAt = now
		_, err = o.Insert(&user)
		if err == nil {
			return &user, 0
		}
		time.Sleep(1 * time.Millisecond)
	}

	beego.Warning("CreateUser fail: ", err)
	return nil, utils.ERROR_CODE_SYSTEM_ERROR
}

func VerifyUserByPhone(phone *string, secret string) (*User, int) {
	utils.AssertNotEmptyString(phone)
	o := orm.NewOrm()
	user := User{Phone: phone}
	err := o.Read(&user, "Phone")

	if err != nil {
		return nil, utils.ERROR_CODE_USERS_USER_NOT_EXISTS
	}

	if utils.Secret2Password("phone:" + *phone, secret) != user.Password {
		return nil, utils.ERROR_CODE_TOKENS_PASSWORD_MISMATCH
	}

	for i := 0; i < DB_UNIQUE_CONFLICT_TRY; i++ {
		user.Token = utils.GenToken()
		user.UpdateAt = utils.GetTimeMillis()
		_, err = o.Update(&user)
		if err == nil {
			return &user, 0
		}
		time.Sleep(1 * time.Millisecond)
	}

	beego.Warning("VerifyUser, update token fail: ", err)
	return nil, utils.ERROR_CODE_SYSTEM_ERROR
}

// TODO return more detailed error number
// NOTE: this will do a fully update, so the user must be read from DB, and then update its field
func UpdateUser(user *User) int {
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		beego.Warning("UpdateUser fail: ", err)
		return utils.ERROR_CODE_SYSTEM_ERROR
	}

	var errNum int
	for i := 0; i < DB_UNIQUE_CONFLICT_TRY; i++ {
		user.UpdateAt = utils.GetTimeMillis()
		errNum = updateUserInternal(&o, user)
		if errNum == 0 {
			break
		}
		time.Sleep(1 * time.Millisecond)
	}
	if errNum > 0 {
		o.Rollback()
		return errNum
	} else {
		err = o.Commit()
		if err != nil {
			beego.Warning("UpdateUser fail: commit fail ", err)
			o.Rollback()
			return utils.ERROR_CODE_SYSTEM_ERROR
		}
	}

	return 0
}

// callee's duty to commit & rollback
func updateUserInternal(o *orm.Ormer, user *User) int {
	_, err := (*o).Update(user)
	if err != nil {
		beego.Warning("UpdateUser, update user fail: ", user, err)
		return utils.ERROR_CODE_SYSTEM_ERROR
	}

	return 0
}
