package http

import (
	"context"
	"io"
	"net"
	"net/http"
	"time"
)

type (
	HTTP struct {
		ConnTimeout, KeepAliveTimeout time.Duration
	}
)

func New() *HTTP {
	var h = new(HTTP)
	h.ConnTimeout = 10 * time.Second
	h.KeepAliveTimeout = 1 * time.Minute
	return h
}

func NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}
func NewClient(connTimeout, keepaliveTimeout time.Duration) http.Client {
	return http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   connTimeout,
				KeepAlive: keepaliveTimeout,
			}).DialContext,
		},
	}
}
func NewUnixClient(sockfile string, connTimeout, keepaliveTimeout time.Duration) http.Client {
	return http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.DialTimeout("unix", sockfile, connTimeout)
			},
		},
	}
	var nd = net.Dialer{}
	nd.KeepAlivea = keepaliveTimeout
	nd.Timeout = connTimeout
	nd.Dial("unix", sockfile)
}

/*
func SetHeaders(req *http.Request, headers map[string]string) {
	var (
		k, v string
	)
	for k, v = range headers {
		req.Header.Set(k, v)
	}
}
func (g *Ghttp) NewClientPool(maxIdleConns int, idleConnTimeout, expectContinueTimeout, connTimeout, keepaliveTimeout time.Duration) http.Client {
	return http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   connTimeout,
				KeepAlive: keepaliveTimeout,
			}).DialContext,
			MaxIdleConns:          maxIdleConns,
			IdleConnTimeout:       idleConnTimeout,
			ExpectContinueTimeout: expectContinueTimeout,
		},
	}
}
func (g *Ghttp) ResponseReadAll() []byte {
	var (
		b   []byte
		err error
	)
	defer g.res.Body.Close()
	if b, err = ioutil.ReadAll(g.res.Body); err != nil {
		return []byte("")
	}
	return b
}
func UnixGet(sockfile, url string, headers map[string]string, body io.Reader) error {
	var (
		req *http.Request
		err error
	)
	if req, err = NewRequest("GET", url, body); err != nil {
		return err
	}
	SetHeaders(headers)
	c := g.NewUnixClient(sockfile)
	if g.res, err = c.Do(g.req); err != nil {
		return err
	}
	return nil
}
func UnixPost(url string, body io.Reader) error {
	var err error
	return err
}
func (g *Ghttp) UnixHead(url string, body io.Reader) error {
	var err error
	return err
}
func (g *Ghttp) Get(url string, headers map[string]string, body io.Reader) error {
	var err error
	if err = g.NewRequest("GET", url, body); err != nil {
		return err
	}
	g.SetHeaders(headers)
	c := g.NewClient(g.connTimeout, g.keepaliveTimeout)
	if g.res, err = c.Do(g.req); err != nil {
		return err
	}
	return nil
}
func (g *Ghttp) Post(url string, body io.Reader) error {
	var err error
	return err
}
*/
