package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"os"
)

func main() {
	remoteAddr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		fmt.Printf("Error resolving address: %v\n", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, remoteAddr)

	re := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		str, err := re.ReadString('\n')
		if err != nil {
			slog.Error(err.Error())
		}
		conn.Write([]byte(str))
	}
}
