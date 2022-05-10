package common

import "reflect"

func TypeOf(err error) string {
	return reflect.TypeOf(err).String()
}
