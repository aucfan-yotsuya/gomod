package ulid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func T_ULID(t *testing.T) {
	assert.Len(t, ULID(), 16, nil)
}
