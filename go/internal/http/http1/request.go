package http1

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"myhttp/internal/http"
	"strconv"
	"strings"
)

type RequestLine struct {
	Method  string
	URI     string
	Version string
}

type Request struct {
	RequestLine   RequestLine
	RequestHeader http.Header
	RequestBody   []byte
}

func (r *Request) Header() http.Header {
	return r.RequestHeader
}

func (r *Request) Body() []byte {
	return r.RequestBody
}

func (r *Request) RequestParams() map[string]string {
	rs := map[string]string{}
	uri := r.RequestLine.URI

	index := strings.LastIndex(uri, "?") + 1
	if index < 0 {
		return rs
	}
	uri = uri[index:]
	pairs := strings.Split(uri, "&")
	for _, s := range pairs {
		fmt.Println(s)
		a := strings.Split(s, "=")
		if len(a) != 2 {
			fmt.Println("Param gì đây?: ", s)
			continue
		}
		rs[a[0]] = a[1]
	}
	return rs
}

func (r *Request) Path() string {
	uri := r.RequestLine.URI
	index := strings.Index(uri, "?")
	if index <= 0 {
		return uri
	}
	return r.RequestLine.URI[:index]
}

var ErrContentLengthHeaderNotFound = errors.New("Content-Length header not found")

func ReadRequest(r *bufio.Reader) (http.Request, error) {
	requestLine, err := readRequestLine(r)
	if err != nil {
		return &Request{}, err
	}
	headers := readHeaders(r)
	if _, ok := headers["Content-Length"]; !ok {
		if requestLine.Method != "GET" {
			return &Request{}, ErrContentLengthHeaderNotFound
		}
		headers["Content-Length"] = []string{"0"}
	}
	contentLength, err := strconv.Atoi(headers["Content-Length"][0])
	if err != nil {
		return &Request{}, err
	}
	body := readBody(r, contentLength)
	return &Request{
		RequestLine:   requestLine,
		RequestHeader: headers,
		RequestBody:   body,
	}, nil
}

// Example:
// POST /upload HTTP/1.1\r\n
func readRequestLine(r *bufio.Reader) (RequestLine, error) {
	requestLine := RequestLine{}

	line, err := r.ReadString('\n')
	if err != nil {
		return RequestLine{}, err
	}
	line = strings.TrimSuffix(line, "\r\n")
	dat := strings.Split(line, " ")
	requestLine.Method = dat[0]
	requestLine.URI = dat[1]
	requestLine.Version = dat[2]
	return requestLine, nil
}

func readHeaders(r *bufio.Reader) http.Header {
	headers := http.Header{}
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		line = strings.TrimSuffix(line, "\r\n")
		if len(line) == 0 {
			break
		}

		dat := strings.Split(line, ": ")
		headers[dat[0]] = append(headers[dat[0]], dat[1])
	}
	return headers
}

func readBody(r *bufio.Reader, contentLength int) []byte {
	if contentLength == 0 {
		return []byte{}
	}
	body := make([]byte, contentLength)
	_, err := r.Read(body)
	if err != nil {
		if err == io.EOF {
			return body
		}
		panic(err)
	}
	return body
}
