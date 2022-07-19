package common

import (
	"math/rand"
	"time"
)

func Shuffle(sl []interface{}) []interface{} {
	rand.Seed(time.Now().UnixNano())
	return rand.Shuffle(len(sl), func(i, j int) { sl[i], sl[j] = sl[j], sl[i] })
}
