package common

import (
	"math/rand"
	"time"
)

func RandomSlice(sl []interface{}) []interface{} {
	var resp []interface{}
	for {
		rand.Seed(time.Now().UnixNano())
		n = rand.Intn(len(sl))
		resp = append(resp, sl[n])
		sl = append(sl[:n], sl[n+1:]...)
		if len(sl) < 1 {
			break
		}
	}
	return resp
}
