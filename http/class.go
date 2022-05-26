package http

import (
	"context"
	"crypto/tls"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"

	"github.com/aucfan-yotsuya/gomod/common"
)

func (h *HTTP) NewClient() *HTTP {
	h.Client = &http.Client{
		Timeout: h.Timeout,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   h.Timeout,
				KeepAlive: h.KeepAliveTimeout,
			}).DialContext,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: h.InsecureSkipVerify,
			},
		},
	}
	return h
}
func (h *HTTP) Do() error {
	if h.NilRequest() {
		return &Err{Message: "Request has nil"}
	}
	h.Response, err = h.Client.Do(h.Request)
	return err
}
func (h *HTTP) Close() {
	if !h.NilRequest() {
		h.Request = nil
	}
	if !h.NilResponse() {
		h.Response.Body.Close()
		h.Response = nil
	}
}
func (h *HTTP) NilRequest() bool  { return h.Request == nil }
func (h *HTTP) NilResponse() bool { return h.Response == nil }
func (h *HTTP) DumpRequest() ([]byte, error) {
	if h.NilRequest() {
		return []byte{}, &Err{Message: "Request has nil"}
	}
	return httputil.DumpRequest(h.Request, h.Debug.Request == DebugAll)
}
func (h *HTTP) DumpResponse() ([]byte, error) {
	if h.NilResponse() {
		return []byte{}, &Err{Message: "Response has nil"}
	}
	return httputil.DumpResponse(h.Response, h.Debug.Response == DebugAll)
}
func (h *HTTP) ResponseHeader() (http.Header, error) {
	if h.NilResponse() {
		return http.Header{}, &Err{Message: "Response has nil"}
	}
	return h.Response.Header, nil
}
func (h *HTTP) ReadAllResponse() ([]byte, error) {
	if h.NilResponse() {
		return []byte{}, &Err{Message: "Response has nil"}
	}
	return ioutil.ReadAll(h.Response.Body)
}
func (h *HTTP) BeforeRequest(opt *Request) error {
	if h.Request, err = http.NewRequestWithContext(
		func() context.Context { ctx, _ := common.Context(h.Timeout); return ctx }(),
		*opt.Method,
		*opt.Url,
		func(i *io.Reader) io.Reader {
			if i == nil {
				return nil
			}
			return *i
		}(opt.Body),
	); err != nil {
		return &Err{Message: err.Error()}
	}
	var k string
	for i, v := range *opt.Header {
		if common.Number(i).Even() {
			k = v
		} else {
			h.Request.Header.Add(k, v)
		}
	}
	return nil
}

// ここまで
// func (h *HTTP) NewClientPool(maxIdleConns int) *HTTP {
// 	h.Client = http.Client{
// 		Transport: &http.Transport{
// 			DialContext: (&net.Dialer{
// 				Timeout:   h.Timeout,
// 				KeepAlive: h.KeepAliveTimeout,
// 			}).DialContext,
// 			MaxIdleConns:          maxIdleConns,
// 			IdleConnTimeout:       h.Timeout,
// 			ExpectContinueTimeout: h.Timeout,
// 		},
// 	}
// 	return h
// }
//func NewUnixClient(sockfile string, connTimeout, keepaliveTimeout time.Duration) *HTTP {
//	h.Client = http.Client{
//		Transport: &http.Transport{
//			DialContext: (&net.Dialer{
//				Timeout:   connTimeout,
//				KeepAlive: keepaliveTimeout,
//			}).DialContext,
//		},
//	}
//	return h
//	/*
//		var nd = net.Dialer{}
//		nd.KeepAlivea = keepaliveTimeout
//		nd.Timeout = connTimeout
//		nd.Dial("unix", sockfile)
//	*/
//}
//func (h *HTTP) SetHeaders(headers map[string]string) *HTTP {
//	if h.NilRequest() {
//		h.Err = &Err{Message: "Request has nil"}
//		return h
//	}
//	for k, v := range headers {
//		h.Request.Header.Add(k, v)
//	}
//	return h
//}
// func (h *HTTP) Do() error {
// 	if h.NilRequest() {
// 		return errors.New("nil http.Request")
// 	}
// 	h.Response, err = h.Client.Do(h.Request)
// 	return err
// }

/*
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
