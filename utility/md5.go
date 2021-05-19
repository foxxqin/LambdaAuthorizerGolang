package utility

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5Hash(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))

	return hex.EncodeToString(hash.Sum(nil))
}
