package helpers

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Md5(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

// EncodeBase64 生成base64（标准方式）
func EncodeBase64(_string string) string {
	res := base64.StdEncoding.EncodeToString([]byte(_string))
	return res
}

// DecodeBase64 解密base64（标准方式）
func DecodeBase64(_string string) string {
	res, err := base64.StdEncoding.DecodeString(_string)
	if err != nil {
		fmt.Printf("DecodeBase64 Error: %s ", err.Error())
		return ""
	}
	return string(res)
}

// EncodeUrlBase64 加密文件和url名安全型base64
func EncodeUrlBase64(_string string) string {
	res := base64.URLEncoding.EncodeToString([]byte(_string))
	return res
}

// DecodeUrlBase64 解密文件和url名安全型base64
func DecodeUrlBase64(_string string) string {
	res, err := base64.URLEncoding.DecodeString(_string)
	if err != nil {
		fmt.Printf("DecodeUrlBase64 Error: %s ", err.Error())
		return ""
	}
	return string(res)
}
