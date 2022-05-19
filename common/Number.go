package common

import "strconv"

type (
	Number int
)

func (n Number) String() string {
	return strconv.Itoa(int(n))
}
func (n Number) Even() bool {
	return int(n)%2 == 0
}
func Atoi(i string) int {
	n, _ := strconv.Atoi(i)
	return n
}
