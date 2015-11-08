package models_test

import (
	"time"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/astaxie/beego/orm"

	"github.com/Piasy/BeegoTDDBootStrap/models"
	"github.com/Piasy/BeegoTDDBootStrap/utils"
)

func TestCreateVerification(t *testing.T) {
	initORM()
	o := orm.NewOrm()
	phone := "18801234567"

	// not exist
	verification := models.Verification{Phone: phone}
	err := o.Read(&verification, "Phone")
	assert.NotNil(t, err)
	assert.Equal(t, phone, verification.Phone)
	assert.Empty(t, verification.Id)
	assert.Empty(t, verification.Code)
	assert.Empty(t, verification.Expire)

	// create one
	errNum := models.CreateVerification(phone)
	assert.Equal(t, 0, errNum)
	verification = models.Verification{Phone: phone}
	err = o.Read(&verification, "Phone")
	assert.Nil(t, err)
	assert.Equal(t, phone, verification.Phone)
	assert.True(t, verification.Id > 0)
	assert.True(t, len(verification.Code) == 6)
	now := time.Now().Unix()
	assert.True(t, now + utils.VERIFY_CODE_EXPIRE_IN_SECONDS - 5 < verification.Expire)
	assert.True(t, verification.Expire < now + utils.VERIFY_CODE_EXPIRE_IN_SECONDS + 5)

	// create it again
	errNum = models.CreateVerification(phone)
	assert.Equal(t, 0, errNum)
	another := models.Verification{Phone: phone}
	err = o.Read(&another, "Phone")
	assert.Nil(t, err)
	assert.Equal(t, verification.Id, another.Id)
	assert.Equal(t, phone, another.Phone)
	assert.True(t, another.Id > 0)
	assert.True(t, len(another.Code) == 6)
	now = time.Now().Unix()
	assert.True(t, now + utils.VERIFY_CODE_EXPIRE_IN_SECONDS - 5 < another.Expire)
	assert.True(t, another.Expire < now + utils.VERIFY_CODE_EXPIRE_IN_SECONDS + 5)

	// clean up
	deleteVerification(t, another.Id)

	// not exist after delete
	verification = models.Verification{Phone: phone}
	err = o.Read(&verification, "Phone")
	assert.NotNil(t, err)
}

func TestCheckVerifyCode(t *testing.T) {
	initORM()
	o := orm.NewOrm()
	phone := "18801234567"

	// not exist
	verification := models.Verification{Phone: phone}
	err := o.Read(&verification, "Phone")
	assert.NotNil(t, err)
	assert.Equal(t, phone, verification.Phone)
	assert.Empty(t, verification.Id)
	assert.Empty(t, verification.Code)
	assert.Empty(t, verification.Expire)

	// create one
	errNum := models.CreateVerification(phone)
	assert.Equal(t, 0, errNum)
	verification = models.Verification{Phone: phone}
	err = o.Read(&verification, "Phone")
	assert.Nil(t, err)
	assert.Equal(t, phone, verification.Phone)
	assert.True(t, verification.Id > 0)
	assert.True(t, len(verification.Code) == 6)
	now := time.Now().Unix()
	assert.True(t, now + utils.VERIFY_CODE_EXPIRE_IN_SECONDS - 5 < verification.Expire)
	assert.True(t, verification.Expire < now + utils.VERIFY_CODE_EXPIRE_IN_SECONDS + 5)

	// check wrong code
	errNum = models.CheckVerifyCode(phone, verification.Code + "1")
	assert.Equal(t, utils.ERROR_CODE_VERIFY_CODE_MISMATCH, errNum)

	// check it
	errNum = models.CheckVerifyCode(phone, verification.Code)
	assert.Equal(t, 0, errNum)

	// check again will fail
	errNum = models.CheckVerifyCode(phone, verification.Code)
	assert.Equal(t, utils.ERROR_CODE_VERIFY_CODE_MISMATCH, errNum)

	// check not exist
	errNum = models.CheckVerifyCode("18812345678", verification.Code)
	assert.Equal(t, utils.ERROR_CODE_VERIFY_CODE_MISMATCH, errNum)

	// simulate expire
	verification.Expire = time.Now().Unix() - 100
	_, err = o.Update(&verification)
	assert.Nil(t, err)

	// check should fail
	errNum = models.CheckVerifyCode(phone, verification.Code)
	assert.Equal(t, utils.ERROR_CODE_VERIFY_CODE_MISMATCH, errNum)

	// clean up
	deleteVerification(t, verification.Id)

	// not exist after delete
	verification = models.Verification{Phone: phone}
	err = o.Read(&verification, "Phone")
	assert.NotNil(t, err)
}

func deleteVerification(t *testing.T, id int64) {
	o := orm.NewOrm()
	verification := models.Verification{Id: id}
	_, err := o.Delete(&verification)
	assert.Nil(t, err)
}