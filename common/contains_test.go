package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	assert.True(t, Contains(&[]string{"a", "b", "c"}, Pstring("b")))
	assert.False(t, Contains(&[]string{"a", "b", "c"}, Pstring("d")))
}
