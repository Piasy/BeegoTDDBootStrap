package utils

func AssertNotEmptyString(str *string) {
	if IsEmptyString(str) {
		panic("Assertion fail, string is empty")
	}
}
