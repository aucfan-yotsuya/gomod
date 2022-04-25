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
