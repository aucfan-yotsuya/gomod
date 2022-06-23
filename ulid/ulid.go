package ulid

import (
	"crypto/rand"

	"github.com/oklog/ulid/v2"
)

func ULID() ulid.ULID {
	return ulid.MustNew(ulid.Now(), rand.Reader)
}
