package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Piasy/BeegoTDDBootStrap/utils"
)

func TestIsValidPhone(t *testing.T) {
	assert.True(t, utils.IsValidPhone("18801234567"))
	assert.True(t, utils.IsValidPhone("13398765432"))

	assert.False(t, utils.IsValidPhone("1339876543"))
	assert.False(t, utils.IsValidPhone("11111111111"))
	assert.False(t, utils.IsValidPhone("51532609"))
}
