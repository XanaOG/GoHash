package gohash

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

func MD5(phrase string) string {
	if phrase == "" {
		return ""
	}
	hash := md5.Sum([]byte(phrase))
	return fmt.Sprintf("%x", hash)
}
func SHA256(phrase string) string {
	if phrase == "" {
		return ""
	}
	hash := sha256.Sum256([]byte(phrase))
	return fmt.Sprintf("%x", hash)
}
func SHA512(phrase string) string {
	if phrase == "" {
		return ""
	}
	hash := sha512.Sum512([]byte(phrase))
	return fmt.Sprintf("%x", hash)
}
func Base64(phrase string) string {
	if phrase == "" {
		return ""
	}
	hash := base64.StdEncoding.EncodeToString([]byte(phrase))
	return fmt.Sprintf("%x", hash)
}
