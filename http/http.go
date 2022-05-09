package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/aucfan-yotsuya/gomod/common"
)

var (
	h   *HTTP
	err error
)

func New() *HTTP {
	h = new(HTTP)
	h.Timeout = 60 * time.Second
	return h
}
func Get(h *HTTP, url string, body io.Reader, headers map[string]string) ([]byte, error) {
	h.BeforeRequest(
		common.StringPtr("GET"),
		&url, &body, &headers,
	)
	h.NewContext()
	if h.Request, err = http.NewRequestWithContext(h.ctx, "GET", url, body); err != nil {
		return nil, &ErrRequest{Message: err.Error()}
	}
	if h.Debug.Request {
		var b []byte
		if b, err = h.DumpRequest(); err != nil {
			return nil, &ErrRequest{Message: err.Error()}
		}
		fmt.Fprintf(os.Stderr, string(b))
	}
	for k, v := range headers {
		h.Request.Header.Add(k, v)
	}
	if h.Response, err = h.NewClient().Do(h.Request); err != nil {
		if h.NilResponse() {
			return nil, &ErrRequest{Message: err.Error()}
		} else {
			return nil, &ErrRequest{Message: err.Error(), StatusCode: h.Response.StatusCode}
		}
	}
	if h.Debug.Response {
		var b []byte
		if b, err = h.DumpResponse(); err != nil {
			return nil, &ErrRequest{Message: err.Error()}
		}
		fmt.Fprintf(os.Stderr, string(b))
	}
	return h.ReadAll()
}
