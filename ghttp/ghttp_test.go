package ghttp

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var g ghttp

func TestNew(t *testing.T) {
	assert.ObjectsAreEqual(g.New(), ghttp{})
}
func TestNewRequest(t *testing.T) {
	assert.Nil(t, g.NewRequest("UNKO", "http://aucfan.com", bytes.NewReader([]byte(""))))
}
func TestGet(t *testing.T) {
	assert.Nil(t, g.Get("https://1.1.1.1", map[string]string{
		"User-Agent": "ghttp/0.0",
	}, nil))
	fmt.Printf("%s\n", g.ResponseReadAll())
}
func TestUnixGet(t *testing.T) {
}
