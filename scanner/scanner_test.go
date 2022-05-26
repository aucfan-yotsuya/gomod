package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func T_New(t *testing.T) {
	assert.NotNil(t, New())
}
