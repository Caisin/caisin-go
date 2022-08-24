package strutil

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

func Md5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func ToMd5(data any) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	h := md5.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil)), nil
}
