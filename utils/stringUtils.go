package utils
import (
	"regexp"
	"net/url"
)

func IsEmptyString(str *string) bool {
	return str == nil || *str == ""
}

func AreStringEquals(str1, str2 *string) bool {
	return str1 == str2 || (str1 != nil && str2 != nil && *str1 == *str2)
}

func GetCharCount(input string) (count int) {
	count = 0
	for range input {
		count++
	}
	return count
}

func IsLegalRestrictedStringWithLength(input string, length int) bool {
	if GetCharCount(input) > length {
		return false
	}

	reg := regexp.MustCompile(`^[a-zA-Z0-9'" \p{Han}]+$`)
	return reg.Match([]byte(input))
}

func IsLegalFreeStringWithLength(input string, length int) bool {
	return GetCharCount(input) <= length
}

func UrlEncode(input string) string {
	url, err := url.Parse(input)
	if err != nil {
		panic(err)
	}
	return url.String()
}
