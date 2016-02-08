package utils_test

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/Piasy/HabitsAPI/utils"
)

func TestAuthWithWeiXin(t *testing.T) {
	_, err := utils.AuthWithWeiXin("wx_openid", "wx_token")
	assert.Equal(t, utils.ERROR_CODE_AUTH_WEIXIN_AUTH_FAIL, err)
}

func TestAuthWithWeiBo(t *testing.T) {
	_, err := utils.AuthWithWeiBo("wb_token")
	assert.Equal(t, utils.ERROR_CODE_AUTH_WEIBO_AUTH_FAIL, err)
}

func TestAuthWithQQ(t *testing.T) {
	_, err := utils.AuthWithQQ("qq_openid", "qq_token", "qq_app_id")
	assert.Equal(t, utils.ERROR_CODE_AUTH_QQ_AUTH_FAIL, err)
}
