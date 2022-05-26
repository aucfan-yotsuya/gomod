package db

import (
	"database/sql"
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
	return strings.Compare(common.TypeOf(err), "*db.Err") == 0
}
func IsNoRows(err error) bool {
	return err == sql.ErrNoRows
}
