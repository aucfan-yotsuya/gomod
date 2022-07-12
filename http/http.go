package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/aucfan-yotsuya/gomod/common"
)

const (
	DebugNone int = iota
	DebugHeaderOnly
	DebugAll
)

type (
	HTTP struct {
		Request                   *http.Request
		Response                  *http.Response
		Client                    *http.Client
		Timeout, KeepAliveTimeout time.Duration
		InsecureSkipVerify        bool
		Sockfile                  string
		Debug                     Debug
	}
	Opt struct {
		Url    *string
		Body   *io.Reader
		Header *[]string
	}
	Debug struct {
		Request  int
		Response int
	}
	Request struct {
		Method *string
		Url    *string
		Body   *io.Reader
		Header *[]string
	}
	Response struct {
		StatusCode int
		Header     http.Header
		Body       []byte
	}
)

var (
	h   *HTTP
	err error
)

func New(timeout time.Duration) *HTTP {
	h = new(HTTP)
	h.Timeout = timeout
	return h
}
func GET(h *HTTP, opt *Opt) (Response, error) {
	defer h.Close()
	if err = h.BeforeRequest(&Request{
		Method: common.Pstring("GET"),
		Url:    opt.Url,
		Body:   opt.Body,
		Header: opt.Header,
	}); err != nil {
		return Response{}, err
	}
	if h.Debug.Request > DebugNone {
		var b []byte
		if b, err = h.DumpRequest(); err != nil {
			return Response{}, err
		}
		fmt.Fprintf(os.Stderr, string(b))
	}
	if err = h.NewClient().Do(); err != nil {
		return Response{}, err
	}
	if h.Debug.Response > DebugNone {
		var b []byte
		if b, err = h.DumpResponse(); err != nil {
			return Response{}, err
		}
		fmt.Fprintf(os.Stderr, string(b))
	}

	defer h.Response.Body.Close()
	var response []byte
	if response, err = h.ReadAllResponse(); err != nil {
		return Response{}, err
	}

	return Response{
		StatusCode: h.Response.StatusCode,
		Header:     h.Response.Header,
		Body:       response,
	}, nil
}
func POST(h *HTTP, opt *Opt) (Response, error) {
	defer h.Close()
	if err = h.BeforeRequest(&Request{
		Method: common.Pstring("POST"),
		Url:    opt.Url,
		Body:   opt.Body,
		Header: opt.Header,
	}); err != nil {
		return Response{}, err
	}
	if h.Debug.Request > DebugNone {
		var b []byte
		if b, err = h.DumpRequest(); err != nil {
			return Response{}, err
		}
		fmt.Fprintf(os.Stderr, string(b))
	}
	if err = h.NewClient().Do(); err != nil {
		return Response{}, err
	}
	if h.Debug.Response > DebugNone {
		var b []byte
		if b, err = h.DumpResponse(); err != nil {
			return Response{}, err
		}
		fmt.Fprintf(os.Stderr, string(b))
	}

	defer h.Response.Body.Close()
	var response []byte
	if response, err = h.ReadAllResponse(); err != nil {
		return Response{}, err
	}

	return Response{
		StatusCode: h.Response.StatusCode,
		Header:     h.Response.Header,
		Body:       response,
	}, nil
}
func PUT(h *HTTP, opt *Opt) (Response, error) {
	defer h.Close()
	if err = h.BeforeRequest(&Request{
		Method: common.Pstring("PUT"),
		Url:    opt.Url,
		Body:   opt.Body,
		Header: opt.Header,
	}); err != nil {
		return Response{}, err
	}
	if h.Debug.Request > DebugNone {
		var b []byte
		if b, err = h.DumpRequest(); err != nil {
			return Response{}, err
		}
		fmt.Fprintf(os.Stderr, string(b))
	}
	if err = h.NewClient().Do(); err != nil {
		return Response{}, err
	}
	if h.Debug.Response > DebugNone {
		var b []byte
		if b, err = h.DumpResponse(); err != nil {
			return Response{}, err
		}
		fmt.Fprintf(os.Stderr, string(b))
	}

	defer h.Response.Body.Close()
	var response []byte
	if response, err = h.ReadAllResponse(); err != nil {
		return Response{}, err
	}

	return Response{
		StatusCode: h.Response.StatusCode,
		Header:     h.Response.Header,
		Body:       response,
	}, nil
}
func DELETE(h *HTTP, opt *Opt) (Response, error) {
	defer h.Close()
	if err = h.BeforeRequest(&Request{
		Method: common.Pstring("DELETE"),
		Url:    opt.Url,
		Body:   opt.Body,
		Header: opt.Header,
	}); err != nil {
		return Response{}, err
	}
	if h.Debug.Request > DebugNone {
		var b []byte
		if b, err = h.DumpRequest(); err != nil {
			return Response{}, err
		}
		fmt.Fprintf(os.Stderr, string(b))
	}
	if err = h.NewClient().Do(); err != nil {
		return Response{}, err
	}
	if h.Debug.Response > DebugNone {
		var b []byte
		if b, err = h.DumpResponse(); err != nil {
			return Response{}, err
		}
		fmt.Fprintf(os.Stderr, string(b))
	}

	defer h.Response.Body.Close()
	var response []byte
	if response, err = h.ReadAllResponse(); err != nil {
		return Response{}, err
	}

	return Response{
		StatusCode: h.Response.StatusCode,
		Header:     h.Response.Header,
		Body:       response,
	}, nil
}
