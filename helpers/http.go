package helpers

import "net/url"

// EncodeURL 转义url或转义其他字符
func EncodeURL(_url string) string {
	return url.QueryEscape(_url)
}

// DecodeURL 解义url
func DecodeURL(_url string) (string, error) {
	return url.QueryUnescape(_url)
}

// GetUrlParam 获取url中的参数（非解码）
func GetUrlParam(_url string, _key string) (value string) {
	u, err := url.Parse(_url)
	values := u.Query()
	if err != nil {
		value = ""
	} else {
		value = values.Get(_key)
	}
	return
}
