package ulid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestULID(t *testing.T) {
	assert.Len(t, ULID(), 16, nil)
}
func TestULIDRand(t *testing.T) {
	assert.Len(t, ULIDRand(), 16, nil)
}
