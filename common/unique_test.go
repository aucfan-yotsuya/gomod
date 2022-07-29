package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueStringSlice(t *testing.T) {
	uniq := UniqueStringSlice([]string{"a", "b", "c", "a"})
	assert.Len(t, uniq, 3)
}
