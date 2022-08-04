package strutil

import (
	"bytes"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func IsBlank(str string) bool {
	if len(str) == 0 {
		return true
	}
	for _, r := range []rune(str) {
		if r != ' ' {
			return false
		}
	}
	return true
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}
func IsParamMapSql(str string) bool {
	return strings.Contains(str, "@")
}

func IsEmpty(str string) bool {
	return len(str) == 0
}

func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

func IsAllBlank(str ...string) bool {
	for _, s := range str {
		if IsNotBlank(s) {
			return false
		}
	}
	return true
}

func ContainsAny(str string, subs ...string) bool {
	for _, s := range subs {
		if strings.Contains(str, s) {
			return true
		}
	}
	return false
}

func IsAllNotBlank(str ...string) bool {
	for _, s := range str {
		if IsBlank(s) {
			return false
		}
	}
	return true
}

func IsNotAllBlank(str ...string) bool {
	return IsAllBlank(str...)
}

func RandNum(size int) string {
	if size == 0 {
		return ""
	}
	bf := bytes.Buffer{}
	for i := 0; i < size; i++ {
		bf.WriteString(strconv.Itoa(rand.Intn(9)))
	}
	return bf.String()
}

func IsAnyBlank(str ...string) bool {
	for _, s := range str {
		if IsBlank(s) {
			return true
		}
	}
	return false
}
func IsAnyEmpty(str ...string) bool {
	for _, s := range str {
		if IsEmpty(s) {
			return true
		}
	}
	return false
}
func IsAllEmpty(str ...string) bool {
	for _, s := range str {
		if IsNotEmpty(s) {
			return false
		}
	}
	return true
}

func IsAllNotEmpty(str ...string) bool {
	for _, s := range str {
		if IsEmpty(s) {
			return false
		}
	}
	return true
}

// Rep 将字符串之间字符替换
func Rep(str, start, end, rep string) string {
	strLen := len(str)
	if strLen == 0 {
		return str
	}
	startIdx := strings.Index(str, start) + 1
	endIdx := strings.Index(str, end)
	ret := str[0:startIdx] + rep + str[endIdx:strLen]
	return ret
}

// Sub 将字符串之间字符截取
func Sub(str, start, end string) string {
	if IsBlank(str) {
		return str
	}
	startIdx := strings.Index(str, start) + len(start)
	endIdx := strings.Index(str, end)
	if endIdx <= startIdx {
		return ""
	}
	ret := str[startIdx:endIdx]
	return ret
}

var (
	spcReg, _   = regexp.Compile("[^ 0-9a-zA-Z\u4e00-\u9fa5]")
	phoneReg, _ = regexp.Compile("\\d{11}")
	numReg, _   = regexp.Compile("^[0-9]*$")
)

// HasSpChapter 是否包含特殊字符
func HasSpChapter(str string) bool {
	return spcReg.MatchString(str)
}

// HasPhone 是否有手机号
func HasPhone(str string) bool {
	return phoneReg.MatchString(str)
}

// IsSameUser 判断是否同一个用户
func IsSameUser(name string) bool {
	if utf8.RuneCountInString(name) >= 4 {
		return HasSpChapter(name) || HasPhone(name)
	}
	return false
}

// IsNum 判断是否是数字
func IsNum(str string) bool {
	return numReg.MatchString(str)
}

// Camel2Case 驼峰式写法转为下划线写法
func Camel2Case(name string) string {
	buffer := bytes.Buffer{}
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.WriteRune('_')
			}
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

// Case2Camel 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// UcFirst 首字母大写
func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// LcFirst 首字母小写
func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
