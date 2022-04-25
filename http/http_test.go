package http

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	req = new(http.Request)
	res = new(http.Response)
	err error
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New())
}
func TestNewRequest(t *testing.T) {
	req, err = NewRequest("UNKO", "http://aucfan.com", bytes.NewReader([]byte("")))
	assert.Nil(t, err)
	assert.ObjectsAreEqual(req, http.Request{})
}
func TestNewClient(t *testing.T) {
	var c = NewClient(10*time.Second, 60*time.Second)
	assert.ObjectsAreEqual(c, http.Client{})
}

// func TestGet(t *testing.T) {
// 	assert.Nil(t, Get("https://1.1.1.1", map[string]string{
// 		"User-Agent": "ghttp/0.0",
// 	}, nil))
// 	fmt.Printf("%s\n", ResponseReadAll())
// }
// func TestUnixGet(t *testing.T) {
// }
