package hello

import "fmt"

func Print() {
	fmt.Println(String())
}
func String() string {
	return "hello"
}
func Bytes() []byte {
	return []byte(String())
}
