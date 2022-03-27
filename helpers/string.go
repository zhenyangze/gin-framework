package helpers

import (
	"html"
	"math/rand"
	"strings"
)

func ReplaceString(text string, _old string, _new string) string {
	if len(text) == 0 {
		return ""
	}
	if len(_old) == 0 {
		return text
	}
	if len(_new) == 0 {
		_new = ""
	}
	text = strings.Replace(text, _old, _new, -1)
	return text
}

func ReplaceRangeString(text string, _start int, _end int, _new string) string {
	if len(text) <= _end {
		_end = len(text) - 1
	}
	if len(_new) == 0 {
		_new = "**"
	}
	return text[:_start] + _new + text[_end:]
}

// ShuffleArray 打乱数组(字符串型数组)
func ShuffleArray(strings []string) string {
	for i := len(strings) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	str := ""
	for i := 0; i < len(strings); i++ {
		str += strings[i]
	}
	return str
}

// FilterInput 过滤Input输入的值
// 转义%、"、'、(、)、!、/、^、*、.、
func FilterInput(_value string) string {
	value := _value

	blackArray := [...]string{ // 这些符号将被转义
		"%", "(", ")", "!", "/", "^", "*", ".", "|", "=",
	}
	changArray := [...]string{ // 这些符号将被替代
		"select", "delete", "char", "insert", "count", "exec", "declare", "update",
	}

	for i := 0; i < len(blackArray); i++ {
		txt := blackArray[i]
		value = ReplaceString(value, txt, EncodeURL(txt))
	}
	for j := 0; j < len(changArray); j++ {
		txt := changArray[j]
		value = ReplaceString(value, txt, "_"+txt)
	}

	value = html.EscapeString(value) // xss

	return value
}

// HideStringValue 隐藏/替换字符串中的某些字符
// 如隐藏手机号：185*******6，调用Common.HideStringValue("18512345506", 3, 10, "*")
func HideStringValue(_string string, start int, end int, replaceValue string) string {
	if len(replaceValue) == 0 {
		replaceValue = "*"
	}
	blackString := _string[start:end]
	replace := ""
	for i := 0; i < len(blackString); i++ {
		replace = replace + replaceValue
	}
	return strings.Replace(_string, blackString, replace, -1)
}
