package http1

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

var (
	simplePostRequest = "POST /upload HTTP/1.1\r\nHost: localhost:8080\r\nContent-Type: application/json\r\nContent-Length: 16\r\n\r\n{\"name\": \"test\"}"
	simpleGetRequest  = "GET /greet?name=floyd HTTP/1.1\r\nHost: localhost:8080\r\nContent-Length:0"
)

func TestSendRequest(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/greeting", "application/json", strings.NewReader(`{"name": "floyd"}`))
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("expecteddd no error, got %v", err)
	}
	fmt.Println(string(body))
}

func TestArbitrary(t *testing.T) {
	input := `haha\n`
	input = strings.TrimSuffix(input, `\n`)
	fmt.Println(input)
	fmt.Println(len(input))
}

func TestParseURI(t *testing.T) {
	req, err := ReadRequest(bufio.NewReader(strings.NewReader(simpleGetRequest)))
	if err != nil {
		panic(err)
	}
	fmt.Println(req.Path())
}
