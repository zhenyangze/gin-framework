package helpers

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
