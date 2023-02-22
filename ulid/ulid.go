package ulid

import (
	"crypto/rand"
	simplerand "math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func ULID() ulid.ULID {
	t := time.Now()
	entropy := ulid.Monotonic(simplerand.New(simplerand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy)
}
func ULIDRand() ulid.ULID {
	return ulid.MustNew(ulid.Now(), rand.Reader)
}
