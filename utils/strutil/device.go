package strutil

import "strings"

var (
	unValidMap = map[string]bool{
		"0":                                    true,
		"9f89c84a559f573636a47ff8daed0d33":     true,
		"00000000-0000-0000-0000-000000000000": true,
		"00000000000000000000000000000000":     true,
		"cd9e459ea708a948d5c2f5a6ca8838cf":     true,
	}
)

func IsValidString(str string) bool {
	if IsBlank(str) {
		return false
	}
	if strings.HasPrefix(str, "__") {
		//快手的无效字符
		return false
	}
	return !unValidMap[str]
}
