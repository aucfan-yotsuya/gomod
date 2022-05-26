package redis

import (
	"strings"

	"github.com/aucfan-yotsuya/gomod/common"
)

type (
	Err struct {
		Message string
	}
)

func (e *Err) Error() string {
	return e.Message
}
func IsErr(err error) bool {
	return strings.Compare(common.TypeOf(err), "*redis.Err") == 0
}
