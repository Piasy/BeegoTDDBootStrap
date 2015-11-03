package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type User struct {
	Id       int64
	Uid      int64
	Phone    string
	Username string
	Password string
	Nickname string
}

func UserExists(u User) bool {
	o := orm.NewOrm()
	user := User{Phone: u.Phone}
	err := o.Read(&user, "Phone")
	return err == nil
}

func GetUser(phone string) (*User, error) {
	o := orm.NewOrm()
	user := User{Phone: phone}
	err := o.Read(&user, "Phone")
	if err == nil {
		beego.Debug("models/User::GetUser() no error")
		return &user, nil
	}
	beego.Debug("models/User::GetUser() error: ", err)
	return nil, errors.New("User not exists")
}