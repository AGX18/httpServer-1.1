package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net"
	"os"
	"strings"
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
		lines := getLinesChannel(conn)
		for line := range lines {
			fmt.Println(line)
		}
		conn.Close()
		fmt.Println("the connection is closed")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	strCh := make(chan string, 8)

	currentLineContents := ""
	buffer := make([]byte, 8)
	go func() {
		for {
			n, err := f.Read(buffer)
			if err != nil {
				if errors.Is(err, io.EOF) {
					currentLineContents += string(buffer[:n])
					if currentLineContents != "" {
						strCh <- fmt.Sprintf("%s", currentLineContents)
						currentLineContents = ""
					}
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				break
			}
			str := string(buffer[:n])
			parts := strings.Split(str, "\n")
			for i := 0; i < len(parts)-1; i++ {
				strCh <- fmt.Sprintf("%s%s", currentLineContents, parts[i])
				currentLineContents = ""
			}
			currentLineContents += parts[len(parts)-1]
		}
		close(strCh)
	}()

	return strCh
}
