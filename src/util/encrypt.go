package util

import (
	"crypto/md5"
	"encoding/hex"
)

// str 需要加密的字符串
// return 加密后的字符串
func Encrypt(str string) string {
	h := md5.New()
	h.Write([]byte(str + "user"))
	encodeByte := h.Sum(nil)
	return hex.EncodeToString(encodeByte)
}
