package http1

import (
	"bufio"
	"fmt"
	"io"
	"myhttp/internal/http"
	"net"
)

type Server struct {
	connections []*net.Conn
	handler     http.Handler
}

func (s *Server) Serve(l net.Listener, handler http.Handler) error {
	s.handler = handler
	return s.serve(l)
}

func (s *Server) serve(l net.Listener) error {
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Loi cmnr ne: ", err)
		}
		go func() {
			req, err := ReadRequest(bufio.NewReader(conn))
			if err != nil && err != io.EOF {
				panic(err)
			}
			s.handler.ServeHTTP(&ResponseWriter{conn: conn}, req)
		}()
	}
}
