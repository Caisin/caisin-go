package strutil

import "regexp"

const dateReg = `^[1-9]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$`

func IsDate(str string) bool {
	matched, _ := regexp.MatchString(dateReg, str)
	return matched
}
