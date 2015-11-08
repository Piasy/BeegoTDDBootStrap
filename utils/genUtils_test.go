package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Piasy/BeegoTDDBootStrap/utils"
)

func TestGenToken(t *testing.T) {
	token := utils.GenToken()

	assert.Equal(t, 40, len(token))
}

func TestGenUid(t *testing.T) {
	uid := utils.GenUid()

	assert.True(t, uid >= utils.USER_MIN_UID)
}

func TestGenVerifyCode(t *testing.T) {
	code := utils.GenVerifyCode()

	assert.True(t, len(code) == utils.VERIFY_CODE_LEN)
}