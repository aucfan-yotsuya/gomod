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
func Pint(i int) *int {
	return &i
}
func Puint(i uint) *uint {
	return &i
}
func PintSlice(intSlice []int) *[]int {
	return &intSlice
}
func Preader(reader io.Reader) *io.Reader {
	return &reader
}
