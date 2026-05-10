package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"

	"github.com/AGX18/httpServer-1.1/internal/request"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	addr := ":42069"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("could not listen for connections: %v\n", err)
	}
	defer listener.Close()

	fmt.Printf("Reading data from %s\n", addr)
	fmt.Println("=====================================")
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error(err.Error(), "remote_addr", conn.RemoteAddr())
		}
		fmt.Println("a connection has been accepted")
		req, err := request.RequestFromReader(conn)
		if err != nil {
			logger.Error(err.Error())
		}
		fmt.Printf("Request line: \n- Method: %s\n- Target: %s\n- Version: %s\n",
			req.RequestLine.Method,
			req.RequestLine.RequestTarget,
			req.RequestLine.HttpVersion,
		)
		conn.Close()
		fmt.Println("the connection is closed")
	}
}
