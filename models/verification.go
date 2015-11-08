package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"

	"github.com/Piasy/BeegoTDDBootStrap/utils"
)

type Verification struct {
	Id     int64 `json:"-"`
	Phone  string `json:"phone"`
	Code   string `json:"code"`
	Expire int64 `json:"expire"`
}

func CreateVerification(phone string) int {
	o := orm.NewOrm()
	code := utils.GenVerifyCode()
	// TODO send verify code
	verification := Verification{Phone: phone}
	_, _, err := o.ReadOrCreate(&verification, "Phone")
	if err != nil {
		beego.Warning("CreateVerification, ReadOrCreate fail: ", err)
		return utils.ERROR_CODE_SYSTEM_ERROR
	}

	verification.Code = code
	verification.Expire = time.Now().Unix() + utils.VERIFY_CODE_EXPIRE_IN_SECONDS
	_, err = o.Update(&verification)
	if err != nil {
		beego.Warning("CreateVerification, Update fail: ", err)
		return utils.ERROR_CODE_SYSTEM_ERROR
	}
	return 0
}

func CheckVerifyCode(phone, code string) int {
	o := orm.NewOrm()
	verification := Verification{Phone: phone}
	err := o.Read(&verification, "Phone")
	if err != nil {
		return utils.ERROR_CODE_VERIFY_CODE_MISMATCH
	}

	if verification.Code == code && time.Now().Unix() <= verification.Expire {
		return invalidateVerification(&verification, o)
	}

	if beego.RunMode == "dev" {
		return invalidateVerification(&verification, o)
	}
	return utils.ERROR_CODE_VERIFY_CODE_MISMATCH
}

func invalidateVerification(verification *Verification, o orm.Ormer) int {
	verification.Expire = 0
	_, err := o.Update(verification)
	if err == nil {
		return 0
	} else {
		beego.Warning("CheckVerifyCode, Update expire fail: ", err)
		return utils.ERROR_CODE_SYSTEM_ERROR
	}
}