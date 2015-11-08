package models_test

import (
	"testing"
	"encoding/json"
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
	beego.RunMode = "prod"
	if ormInitiated {
		return
	}
	appConf, err := config.NewConfig("ini", "../conf/app.conf")
	if err != nil {
		panic(err)
	}
	dbUser := appConf.String("admin::dbUser")
	dbPass := appConf.String("admin::dbPass")
	dbName := "beego_unit_test"

	orm.RegisterDriver("mymysql", orm.DR_MySQL)

	var conn string
	if dbPass == "" {
		conn = fmt.Sprintf("%s:@/%s?charset=utf8", dbUser, dbName)
	} else {
		conn = fmt.Sprintf("%s:%s@/%s?charset=utf8", dbUser, dbPass, dbName)
	}
	orm.RegisterDataBase("default", "mysql", conn)
	ormInitiated = true
}

func TestDeSerial(t *testing.T) {
	mock := "{\"nickname\":\"Piasy\",\"phone\":\"18801234567\",\"uid\":1905378617,\"username\":\"wx:piasy_umumu\"}"
	var user models.User
	err := json.Unmarshal([]byte(mock), &user)

	assert.Nil(t, err)

	assert.Equal(t, "Piasy", user.Nickname)
	assert.Equal(t, "18801234567", user.Phone)
	assert.Equal(t, int64(1905378617), user.Uid)
	assert.Equal(t, "wx:piasy_umumu", user.Username)

	assert.Empty(t, user.Id)
	assert.Empty(t, user.Password)
}

func TestSerial(t *testing.T) {
	user := models.User{Id: 1, Nickname: "Piasy", Uid: 1905378617}
	data, err := json.Marshal(user)

	assert.Nil(t, err)

	var fromJson models.User
	err = json.Unmarshal(data, &fromJson)

	assert.Nil(t, err)

	assert.Empty(t, fromJson.Id)
	assert.Empty(t, fromJson.Password)
	assert.Empty(t, fromJson.Username)
	assert.Empty(t, fromJson.Phone)

	assert.Equal(t, user.Nickname, fromJson.Nickname)
	assert.Equal(t, user.Uid, fromJson.Uid)
}

func TestUserPhoneExists(t *testing.T) {
	initORM()
	exists := models.UserPhoneExists("18801234567")
	assert.False(t, exists)
}

func TestGetUserByPhone(t *testing.T) {
	initORM()
	user, err := models.GetUserByPhone("18801234567")
	assert.Nil(t, user)
	assert.Equal(t, utils.ERROR_CODE_USERS_USER_NOT_EXISTS, err)
}

func TestGetUserByUid(t *testing.T) {
	initORM()
	user, err := models.GetUserByUid(2971788563)
	assert.Nil(t, user)
	assert.Equal(t, utils.ERROR_CODE_USERS_USER_NOT_EXISTS, err)
}

func TestGetUserByToken(t *testing.T) {
	initORM()
	user, err := models.GetUserByToken("lgJYnQXrKVPoInPTPnokdPOZISzosxQzNUceRJyA")
	assert.Nil(t, user)
	assert.Equal(t, utils.ERROR_CODE_TOKENS_INVALID_TOKEN, err)
}

func TestCreateUserByPhone(t *testing.T) {
	initORM()

	phone := "18801234567"
	username := "phone:18801234567"
	secret := "8428d916f8cca9ba5971bf58b34d38da20bc3dff"
	password := "6fced8fa30df2eea13ee553d0688089da1d0b81e"

	// no such user before create
	user, err := models.GetUserByPhone(phone)
	assert.Nil(t, user)
	assert.Equal(t, utils.ERROR_CODE_USERS_USER_NOT_EXISTS, err)

	// insert one
	user, err = models.CreateUserByPhone(phone, secret)
	assert.NotNil(t, user)
	assert.Equal(t, 0, err)
	assert.Empty(t, user.Nickname)
	assert.True(t, user.Uid >= utils.USER_MIN_UID)
	assert.Equal(t, password, user.Password)
	assert.Equal(t, username, user.Username)
	assert.True(t, len(user.Token) == 40)
	assert.Equal(t, phone, user.Phone)

	// get it by phone
	getByPhone, err := models.GetUserByPhone(user.Phone)
	assert.Equal(t, 0, err)
	assertUserEquals(t, user, getByPhone)

	// get it by uid
	getByUid, err := models.GetUserByUid(user.Uid)
	assert.Equal(t, 0, err)
	assertUserEquals(t, user, getByUid)

	// get it by token
	getByToken, err := models.GetUserByToken(user.Token)
	assert.Equal(t, 0, err)
	assertUserEquals(t, user, getByToken)

	// clean up
	deleteUser(t, user.Id)

	// no such user after delete
	user, err = models.GetUserByPhone(phone)
	assert.Nil(t, user)
	assert.Equal(t, utils.ERROR_CODE_USERS_USER_NOT_EXISTS, err)
}

func TestVerifyUserByPhone(t *testing.T) {
	initORM()

	phone := "18801234567"
	username := "phone:18801234567"
	secret := "8428d916f8cca9ba5971bf58b34d38da20bc3dff"
	password := "6fced8fa30df2eea13ee553d0688089da1d0b81e"

	// no such user before create
	user, err := models.GetUserByPhone(phone)
	assert.Nil(t, user)
	assert.Equal(t, utils.ERROR_CODE_USERS_USER_NOT_EXISTS, err)


	// insert one
	user, err = models.CreateUserByPhone(phone, secret)
	assert.NotNil(t, user)
	assert.Equal(t, 0, err)
	assert.Empty(t, user.Nickname)
	assert.True(t, user.Uid >= utils.USER_MIN_UID)
	assert.Equal(t, password, user.Password)
	assert.Equal(t, username, user.Username)
	assert.True(t, len(user.Token) == 40)
	assert.Equal(t, phone, user.Phone)

	// get it by phone
	getByPhone, err := models.GetUserByPhone(user.Phone)
	assert.Equal(t, 0, err)
	assertUserEquals(t, user, getByPhone)

	// verify by phone
	verifyByPhone, err := models.VerifyUserByPhone(phone, secret)
	assert.Equal(t, 0, err)
	assertUserEqualsWithoutToken(t, user, verifyByPhone)

	// clean up
	deleteUser(t, user.Id)

	// no such user after delete
	user, err = models.GetUserByPhone(phone)
	assert.Nil(t, user)
	assert.Equal(t, utils.ERROR_CODE_USERS_USER_NOT_EXISTS, err)
}

func assertUserEquals(t *testing.T, expect, actual *models.User) {
	assert.Equal(t, expect.Id, actual.Id)
	assert.Equal(t, expect.Uid, actual.Uid)
	assert.Equal(t, expect.Token, actual.Token)
	assert.Equal(t, expect.Username, actual.Username)
	assert.Equal(t, expect.Phone, actual.Phone)
	assert.Equal(t, expect.Password, actual.Password)
	assert.Equal(t, expect.Nickname, actual.Nickname)
}

func assertUserEqualsWithoutToken(t *testing.T, expect, actual *models.User) {
	assert.Equal(t, expect.Id, actual.Id)
	assert.Equal(t, expect.Uid, actual.Uid)
	assert.NotEqual(t, expect.Token, actual.Token)
	assert.Equal(t, expect.Username, actual.Username)
	assert.Equal(t, expect.Phone, actual.Phone)
	assert.Equal(t, expect.Password, actual.Password)
	assert.Equal(t, expect.Nickname, actual.Nickname)
}

func deleteUser(t *testing.T, id int64) {
	o := orm.NewOrm()
	user := models.User{Id: id}
	_, err := o.Delete(&user)
	assert.Nil(t, err)
}