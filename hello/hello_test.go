package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert.Equal(t, String(), "hello", "TestString")
}
func TestBytes(t *testing.T) {
	assert.Equal(t, Bytes(), []byte("hello"), "TestString")
}
