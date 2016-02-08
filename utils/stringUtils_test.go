package utils_test

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/Piasy/HabitsAPI/utils"
)

func TestIsEmptyString(t *testing.T) {
	assert.True(t, utils.IsEmptyString(nil))
	empty := ""
	assert.True(t, utils.IsEmptyString(&empty))

	abc := "abc"
	assert.False(t, utils.IsEmptyString(&abc))
}

func TestAreStringEquals(t *testing.T) {
	assert.True(t, utils.AreStringEquals(nil, nil))
	empty := ""
	anotherEmpty := ""

	assert.True(t, utils.AreStringEquals(&empty, &empty))
	assert.True(t, utils.AreStringEquals(&empty, &anotherEmpty))
	// won't compile
	//assert.False(t, utils.AreStringEquals(&empty, ""))
	assert.False(t, utils.AreStringEquals(&empty, nil))

	abc := "abc"
	anotherAbc := "abc"
	abcd := "abcd"
	assert.False(t, utils.AreStringEquals(&abc, nil))
	assert.True(t, utils.AreStringEquals(&abc, &abc))
	assert.True(t, utils.AreStringEquals(&abc, &anotherAbc))
	assert.False(t, utils.AreStringEquals(&abc, &empty))
	assert.False(t, utils.AreStringEquals(&abc, &abcd))
}

func TestGetCharCount(t *testing.T) {
	assert.Equal(t, 8, utils.GetCharCount("abc123你好"))
	assert.Equal(t, 9, utils.GetCharCount("abc1234你好"))
	assert.Equal(t, 8, utils.GetCharCount("abc12@你好"))
	assert.Equal(t, 4, utils.GetCharCount("abc\ue050"))
	assert.Equal(t, 11, utils.GetCharCount("Test album1"))
}

func TestIsIllegalRestrictedStringWithLength(t *testing.T) {
	assert.True(t, utils.IsLegalRestrictedStringWithLength("abc123你好", 8))
	assert.True(t, utils.IsLegalRestrictedStringWithLength("\"", 20))

	assert.False(t, utils.IsLegalRestrictedStringWithLength("，", 20))
	assert.False(t, utils.IsLegalRestrictedStringWithLength("’", 20))
	assert.False(t, utils.IsLegalRestrictedStringWithLength("、", 20))
	assert.False(t, utils.IsLegalRestrictedStringWithLength("：", 20))
	assert.False(t, utils.IsLegalRestrictedStringWithLength("；", 20))
	assert.False(t, utils.IsLegalRestrictedStringWithLength("“", 20))
	assert.False(t, utils.IsLegalRestrictedStringWithLength("a1\"你好，。’、：；“", 20))

	assert.False(t, utils.IsLegalRestrictedStringWithLength("abc1234你好", 8))
	assert.False(t, utils.IsLegalRestrictedStringWithLength("abc12@你好", 8))
	assert.False(t, utils.IsLegalRestrictedStringWithLength("abc\ue050", 8))

	assert.True(t, utils.IsLegalRestrictedStringWithLength("Test album1", 12))
	assert.True(t, utils.IsLegalRestrictedStringWithLength("I'm Piasy", 20))
}

func TestIsIllegalFreeStringWithLength(t *testing.T) {
	assert.True(t, utils.IsLegalFreeStringWithLength("abc123你好", 8))
	assert.True(t, utils.IsLegalFreeStringWithLength("\"", 20))

	assert.True(t, utils.IsLegalFreeStringWithLength("，", 20))
	assert.True(t, utils.IsLegalFreeStringWithLength("’", 20))
	assert.True(t, utils.IsLegalFreeStringWithLength("、", 20))
	assert.True(t, utils.IsLegalFreeStringWithLength("：", 20))
	assert.True(t, utils.IsLegalFreeStringWithLength("；", 20))
	assert.True(t, utils.IsLegalFreeStringWithLength("“", 20))
	assert.True(t, utils.IsLegalFreeStringWithLength("a1\"你好，。’、：；“", 20))

	assert.False(t, utils.IsLegalFreeStringWithLength("abc1234你好", 8))
	assert.True(t, utils.IsLegalFreeStringWithLength("abc12@你好", 8))
	assert.True(t, utils.IsLegalFreeStringWithLength("abc\ue050", 8))

	assert.True(t, utils.IsLegalFreeStringWithLength("Test album1", 12))
	assert.True(t, utils.IsLegalFreeStringWithLength("I'm Piasy", 20))
}

func TestUrlEncode(t *testing.T) {
	assert.Equal(t,
		"http://masterpieces.oss-cn-hangzhou.aliyuncs.com/masterpieces/1845099618/1/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202015-11-16%20%E4%B8%8B%E5%8D%882.22.09.png",
		utils.UrlEncode("http://masterpieces.oss-cn-hangzhou.aliyuncs.com/masterpieces/1845099618/1/屏幕快照 2015-11-16 下午2.22.09.png"))
}
