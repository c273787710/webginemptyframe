package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Encry(val string)string{
	m := md5.New()
	m.Write([]byte(val))
	return hex.EncodeToString(m.Sum(nil))
}