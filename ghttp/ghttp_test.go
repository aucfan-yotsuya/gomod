package ghttp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.ObjectsAreEqual(New(), Ghttp{})
}

// func TestNewRequest(t *testing.T) {
// 	assert.Nil(t, NewRequest("UNKO", "http://aucfan.com", bytes.NewReader([]byte(""))))
// }
// func TestGet(t *testing.T) {
// 	assert.Nil(t, Get("https://1.1.1.1", map[string]string{
// 		"User-Agent": "ghttp/0.0",
// 	}, nil))
// 	fmt.Printf("%s\n", ResponseReadAll())
// }
// func TestUnixGet(t *testing.T) {
// }
