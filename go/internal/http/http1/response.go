package http1

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

type Response struct {
	StatusCode int
	Status     string
	Headers    map[string]string
	Body       []byte
}

type ResponseWriter struct {
	conn net.Conn
}

func (w *ResponseWriter) Write(data []byte) (int, error) {
	response := &Response{
		Body:       data,
		StatusCode: 200,
		Status:     "OK",
		Headers:    map[string]string{},
	}
	if json.Valid(data) {
		response.Headers["Content-Type"] = "application/json"
	} else {
		// TODO: support more type
		response.Headers["Content-Type"] = "text/html; charset=utf-8"
	}
	return WriteResponse(w.conn, response)
}

func WriteResponse(w io.Writer, response *Response) (int, error) {
	headers := make(map[string]string)
	for key, value := range response.Headers {
		headers[key] = value
	}
	if len(response.Body) > 0 {
		headers["Content-Length"] = strconv.Itoa(len(response.Body))
	}
	if _, ok := headers["Content-Type"]; !ok {
		headers["Content-Type"] = "application/json"
	}
	if _, ok := headers["Date"]; !ok {
		headers["Date"] = time.Now().Format(time.RFC1123)
	}
	if _, ok := headers["Access-Control-Allow-Origin"]; !ok {
		headers["Access-Control-Allow-Origin"] = "*"
	}
	strResponse := fmt.Sprintf("HTTP/1.1 %d %s\r\n", response.StatusCode, response.Status)
	for key, value := range headers {
		strResponse += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	strResponse += "\r\n"
	strResponse += string(response.Body)

	_, err := w.Write([]byte(strResponse))
	return len(response.Body), err
}
