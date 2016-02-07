package models

import (
	"github.com/astaxie/beego/orm"

	"github.com/Piasy/BeegoTDDBootStrap/utils"
	"github.com/astaxie/beego"
)

type User struct {
	Id       int64 `json:"-"`
	Uid      int64 `json:"uid"`
	Username string `json:"username,omitempty"`
	Password string `json:"-"`
	Nickname string `json:"nickname"`
	Token    string `json:"token,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

const USERS_TABLE_NAME string = "users"

func (u *User) TableName() string {
	return USERS_TABLE_NAME
}

func UserPhoneExists(phone string) bool {
	o := orm.NewOrm()
	user := User{Username: "phone:" + phone}
	err := o.Read(&user, "Username")
	return err == nil
}

func GetUserByPhone(phone string) (*User, int) {
	o := orm.NewOrm()
	user := User{Phone: phone}
	err := o.Read(&user, "Phone")
	if err == nil {
		return &user, 0
	}
	return nil, utils.ERROR_CODE_USERS_USER_NOT_EXISTS
}

func GetUserByUid(uid int64) (*User, int) {
	o := orm.NewOrm()
	user := User{Uid: uid}
	err := o.Read(&user, "Uid")
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

func CreateUserByPhone(phone, secret string) (*User, int) {
	o := orm.NewOrm()
	password := utils.Secret2Password("phone:" + phone, secret)
	user := User{Username: "phone:" + phone, Password: password, Phone: phone}

	var err error
	for i := 0; i < DB_UNIQUE_CONFLICT_TRY; i++ {
		user.Uid = utils.GenUid()
		user.Token = utils.GenToken()
		_, err = o.Insert(&user)
		if err == nil {
			return &user, 0
		}
	}

	beego.Warning("CreateUser fail: ", err)
	return nil, utils.ERROR_CODE_SYSTEM_ERROR
}

func VerifyUserByPhone(phone, secret string) (*User, int) {
	o := orm.NewOrm()
	user := User{Username: "phone:" + phone}
	err := o.Read(&user, "Username")

	if err != nil {
		return nil, utils.ERROR_CODE_USERS_USER_NOT_EXISTS
	}

	if utils.Secret2Password("phone:" + phone, secret) != user.Password {
		return nil, utils.ERROR_CODE_TOKENS_PASSWORD_MISMATCH
	}

	for i := 0; i < DB_UNIQUE_CONFLICT_TRY; i++ {
		user.Token = utils.GenToken()
		_, err = o.Update(&user)
		if err == nil {
			return &user, 0
		}
	}

	beego.Warning("VerifyUser, update token fail: ", err)
	return nil, utils.ERROR_CODE_SYSTEM_ERROR
}
