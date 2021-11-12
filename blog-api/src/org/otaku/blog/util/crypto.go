package util

import (
	"crypto/sha1"
	"fmt"
)


func Sha1(content string, solt string) (string, error) {
	toHash := content + solt
	h := sha1.New()
	if _, err := h.Write([]byte(toHash)); err != nil {
		return "", err
	}
	hashBytes := h.Sum(nil)
	//转成16进制
	return fmt.Sprintf("%x", hashBytes), nil
}
