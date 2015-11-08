package utils

import (
	"github.com/ttacon/libphonenumber"
)

func IsValidPhone(phone string) bool {
	num, err := libphonenumber.Parse(phone, "CN")
	return err == nil && num != nil && libphonenumber.GetNumberType(num) == libphonenumber.MOBILE
}