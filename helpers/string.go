package helpers

import (
	"html"
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
// 如隐藏手机号：185*******6，调用Common.HideStringValue("18511111111", 3, 10, "*")
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

func UcFirst(_string string) string {
	if len(_string) == 0 {
		return ""
	}

	s := []rune(_string)
	if s[0] >= 97 && s[0] <= 122 {
		s[0] -= 32
	}
	return string(s)
}

func LcFirst(_string string) string {
	if len(_string) == 0 {
		return ""
	}

	s := []rune(_string)
	if s[0] >= 65 && s[0] <= 90 {
		s[0] += 32
	}
	return string(s)
}

func Trim(str, character_mask string) string {

	if character_mask == "" {
		character_mask = " \r\n\t\x0B"
	}

	return strings.Trim(str, character_mask)
}

func Ltrim(str, character_mask string) string {

	if character_mask == "" {
		character_mask = " \r\n\t\x0B"
	}

	return strings.TrimLeft(str, character_mask)
}

func Rtrim(str, character_mask string) string {

	if character_mask == "" {
		character_mask = " \r\n\t\x0B"
	}

	return strings.TrimRight(str, character_mask)
}
