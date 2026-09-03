package http

import (
	"net"
)

type Header map[string][]string

type Request interface {
	Header() Header
	Body() []byte
	RequestParams() map[string]string
	Path() string
}

type ResponseWriter interface {
	Write([]byte) (int, error)

	// Header() Header
	// WriteHeader(statusCode int)
}

type Handler interface {
	ServeHTTP(ResponseWriter, Request)
}

type HttpServer interface {
	Serve(l net.Listener, handler Handler) error
}
