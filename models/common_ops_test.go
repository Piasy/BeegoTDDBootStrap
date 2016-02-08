package models_test

import (
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"

	"github.com/Piasy/BeegoTDDBootStrap/models"
	"github.com/Piasy/BeegoTDDBootStrap/utils"
)

var ormInitiated bool = false

func initORM() {
	// switch to prod
	beego.BConfig.RunMode = "prod"
	if ormInitiated {
		return
	}
	appConf, err := config.NewConfig("ini", "../conf/app.conf")
	if err != nil {
		panic(err)
	}
	dbAddr := appConf.String("admin::dbAddr")
	dbUser := appConf.String("admin::dbUser")
	dbPass := appConf.String("admin::dbPass")

	orm.RegisterDriver("mymysql", orm.DRMySQL)
	conn := fmt.Sprintf("%s:%s@tcp(%s)/beego_unit_test?charset=utf8mb4", dbUser, dbPass, dbAddr)
	orm.RegisterDataBase("default", "mysql", conn)
	ormInitiated = true
}

func deleteVerification(t *testing.T, id int64) {
	o := orm.NewOrm()
	verification := models.Verification{Id: id}
	_, err := o.Delete(&verification)
	assert.Nil(t, err)
}

func assertUserEquals(t *testing.T, expect, actual *models.User) {
	assert.Equal(t, expect.Id, actual.Id)

	assert.Equal(t, expect.Uid, actual.Uid)
	assert.Equal(t, expect.Token, actual.Token)
	assert.True(t, utils.AreStringEquals(actual.Phone, expect.Phone))
	assert.True(t, utils.AreStringEquals(actual.WeiXin, expect.WeiXin))
	assert.True(t, utils.AreStringEquals(actual.WeiBo, expect.WeiBo))
	assert.True(t, utils.AreStringEquals(actual.QQ, expect.QQ))

	assert.Equal(t, expect.Password, actual.Password)
	assert.Equal(t, expect.Nickname, actual.Nickname)
	assert.Equal(t, expect.QQNickName, actual.QQNickName)
	assert.Equal(t, expect.WeiBoNickName, actual.WeiBoNickName)
	assert.Equal(t, expect.WeiXinNickName, actual.WeiXinNickName)
	assert.Equal(t, expect.Gender, actual.Gender)
	assert.Equal(t, expect.Avatar, actual.Avatar)

	assert.Equal(t, expect.CreateAt, actual.CreateAt)
	assert.Equal(t, expect.UpdateAt, actual.UpdateAt)

}

func assertUserEqualsWithoutToken(t *testing.T, expect, actual *models.User) {
	assert.Equal(t, expect.Id, actual.Id)

	assert.Equal(t, expect.Uid, actual.Uid)
	assert.True(t, utils.AreStringEquals(actual.Phone, expect.Phone))
	assert.True(t, utils.AreStringEquals(actual.WeiXin, expect.WeiXin))
	assert.True(t, utils.AreStringEquals(actual.WeiBo, expect.WeiBo))
	assert.True(t, utils.AreStringEquals(actual.QQ, expect.QQ))

	assert.Equal(t, expect.Password, actual.Password)
	assert.Equal(t, expect.Nickname, actual.Nickname)
	assert.Equal(t, expect.QQNickName, actual.QQNickName)
	assert.Equal(t, expect.WeiBoNickName, actual.WeiBoNickName)
	assert.Equal(t, expect.WeiXinNickName, actual.WeiXinNickName)
	assert.Equal(t, expect.Gender, actual.Gender)
	assert.Equal(t, expect.Avatar, actual.Avatar)

	assert.Equal(t, expect.CreateAt, actual.CreateAt)
	assert.Equal(t, expect.UpdateAt, actual.UpdateAt)
}

func deleteUser(t *testing.T, id int64) {
	o := orm.NewOrm()
	user := models.User{Id: id}
	_, err := o.Delete(&user)
	assert.Nil(t, err)
}
