package common

import "io"

var (
	err error
)

func Pstring(str string) *string {
	return &str
}
func PstringSlice(strSlice []string) *[]string {
	return &strSlice
}
func Preader(reader io.Reader) *io.Reader {
	return &reader
}
