package models_test

import (
	"testing"
	"encoding/json"

	"github.com/stretchr/testify/assert"
	"github.com/astaxie/beego/orm"

	"github.com/Piasy/BeegoTDDBootStrap/models"
	"github.com/Piasy/BeegoTDDBootStrap/utils"
)

func TestUserDeSerial(t *testing.T) {
	mock := "{\"nickname\":\"Piasy\",\"phone\":\"18801234567\",\"uid\":1905378617}"
	phone := "18801234567"
	var user models.User
	err := json.Unmarshal([]byte(mock), &user)

	assert.Nil(t, err)

	assert.Equal(t, "Piasy", user.Nickname)
	assert.True(t, utils.AreStringEquals(user.Phone, &phone))
	assert.Equal(t, int64(1905378617), user.Uid)

	assert.Empty(t, user.Id)
	assert.Empty(t, user.Password)
}

func TestUserSerial(t *testing.T) {
	user := models.User{Id: 1, Nickname: "Piasy", Uid: 1905378617}
	data, err := json.Marshal(user)

	assert.Nil(t, err)

	var fromJson models.User
	err = json.Unmarshal(data, &fromJson)

	assert.Nil(t, err)

	assert.Empty(t, fromJson.Id)
	assert.Empty(t, fromJson.Password)
	assert.True(t, utils.IsEmptyString(fromJson.Phone))

	assert.Equal(t, user.Nickname, fromJson.Nickname)
	assert.Equal(t, user.Uid, fromJson.Uid)
}

func TestUserPhoneExists(t *testing.T) {
	initORM()
	phone := "18801234567"
	exists := models.UserPhoneExists(&phone)
	assert.False(t, exists)
}

func TestGetUserByUidNotExist(t *testing.T) {
	initORM()
	user, err := models.GetUserByUid(2971788563)
	assert.Nil(t, user)
	assert.Equal(t, utils.ERROR_CODE_USERS_USER_NOT_EXISTS, err)
}

func TestGetUserByUid(t *testing.T) {
	initORM()
	user := models.User{Uid: 2971788563, Nickname: "Piasy", Gender: 1}
	o := orm.NewOrm()
	id, err := o.Insert(&user)
	assert.Equal(t, user.Id, id)
	assert.Nil(t, err)
	got, errNum := models.GetUserByUid(user.Uid)
	assert.Zero(t, errNum)
	assertUserEquals(t, &user, got)

	deleteUser(t, user.Id)
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
	secret := "8428d916f8cca9ba5971bf58b34d38da20bc3dff"
	password := "6fced8fa30df2eea13ee553d0688089da1d0b81e"

	// insert one
	user, err := models.CreateUserByPhone(&phone, secret)
	assert.NotNil(t, user)
	assert.Zero(t, err)
	assert.Empty(t, user.Nickname)
	assert.True(t, user.Uid >= utils.USER_MIN_UID)
	assert.Equal(t, password, user.Password)
	assert.True(t, len(user.Token) == 40)
	assert.True(t, utils.AreStringEquals(user.Phone, &phone))
	now := utils.GetTimeMillis()
	assert.True(t, now - 1000 < user.CreateAt)
	assert.True(t, user.CreateAt < now + 1000)

	// get it by phone
	getByPhone, err := models.GetUserByUid(user.Uid)
	assert.Zero(t, err)
	assertUserEquals(t, user, getByPhone)

	// get it by uid
	getByUid, err := models.GetUserByUid(user.Uid)
	assert.Zero(t, err)
	assertUserEquals(t, user, getByUid)

	// get it by token
	getByToken, err := models.GetUserByToken(user.Token)
	assert.Zero(t, err)
	assertUserEquals(t, user, getByToken)

	// clean up
	deleteUser(t, user.Id)

	// no such user after delete
	user, err = models.GetUserByUid(user.Uid)
	assert.Nil(t, user)
	assert.Equal(t, utils.ERROR_CODE_USERS_USER_NOT_EXISTS, err)
}

func TestVerifyUserByPhone(t *testing.T) {
	initORM()

	phone := "18801234567"
	secret := "8428d916f8cca9ba5971bf58b34d38da20bc3dff"
	password := "6fced8fa30df2eea13ee553d0688089da1d0b81e"

	// insert one
	user, err := models.CreateUserByPhone(&phone, secret)
	assert.NotNil(t, user)
	assert.Zero(t, err)
	assert.Empty(t, user.Nickname)
	assert.True(t, user.Uid >= utils.USER_MIN_UID)
	assert.Equal(t, password, user.Password)
	assert.True(t, len(user.Token) == 40)
	assert.True(t, utils.AreStringEquals(user.Phone, &phone))

	// get it by phone
	getByPhone, err := models.GetUserByUid(user.Uid)
	assert.Zero(t, err)
	assertUserEquals(t, user, getByPhone)

	// verify by phone
	verifyByPhone, err := models.VerifyUserByPhone(&phone, secret)
	assert.Zero(t, err)
	user.UpdateAt = verifyByPhone.UpdateAt
	assertUserEqualsWithoutToken(t, user, verifyByPhone)

	// clean up
	deleteUser(t, user.Id)

	// no such user after delete
	user, err = models.GetUserByUid(user.Uid)
	assert.Nil(t, user)
	assert.Equal(t, utils.ERROR_CODE_USERS_USER_NOT_EXISTS, err)
}

func TestUpdateUser(t *testing.T) {
	phone := "18801234567"
	secret := "8428d916f8cca9ba5971bf58b34d38da20bc3dff"
	phone2 := "18801234568"

	// insert two
	user, err := models.CreateUserByPhone(&phone, secret)
	assert.Zero(t, err)
	user2, err := models.CreateUserByPhone(&phone2, secret)
	assert.Zero(t, err)

	weixin := "wx:piasy_umumu"
	user.WeiXin = &weixin
	user.Nickname = "Piasy"
	user.Gender = 1
	err = models.UpdateUser(user)
	assert.Zero(t, err)

	got, err := models.GetUserByToken(user.Token)
	assert.Zero(t, err)
	assertUserEquals(t, user, got)

	user.Phone = &phone2
	err = models.UpdateUser(user)
	assert.Equal(t, utils.ERROR_CODE_SYSTEM_ERROR, err)

	deleteUser(t, user.Id)
	deleteUser(t, user2.Id)
}

func TestMysqlLikeSearch(t *testing.T) {
	initORM()
	o := orm.NewOrm()
	user1 := models.User{Uid: utils.GenUid(), Token: utils.GenToken(),
		CreateAt: utils.GetTimeMillis(), Nickname: "张三李四", UpdateAt: utils.GetTimeMillis() + 5}
	_, err := o.Insert(&user1)
	assert.Nil(t, err)
	user2 := models.User{Uid: utils.GenUid(), Token: utils.GenToken(),
		CreateAt: utils.GetTimeMillis() + 10, Nickname: "张李三四", UpdateAt: utils.GetTimeMillis() + 15}
	_, err = o.Insert(&user2)
	assert.Nil(t, err)

	var users *[]models.User
	users = new([]models.User)

	num, err := o.QueryTable(new(models.User)).Filter("Nickname__contains", "张三").All(users)

	assert.Nil(t, err)
	assert.Equal(t, int64(1), num)
	assertUserEquals(t, &user1, &((*users)[0]))

	deleteUser(t, user1.Id)
	deleteUser(t, user2.Id)
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
