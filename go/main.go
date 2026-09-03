package main

import (
	"fmt"
	"myhttp/internal/http/http1"
	"net"
)

func main() {
	port := ":8080"
	l, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Listening on %s\n", port)
	server := &http1.Server{}
	server.Serve(l, &myHandler{})
}
