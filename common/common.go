package common

var (
	err error
)

func StringPtr(str string) *string {
	return &str
}
