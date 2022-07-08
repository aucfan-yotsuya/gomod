package common

import (
	"crypto/md5"
	"fmt"
)

func Md5Sum(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
