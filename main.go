package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fd, err := os.Open("messages.txt")
	if err != nil {
		fmt.Printf("couldn't read the file")
		return
	}
	buf := make([]byte, 8)
	currLine := ""
	for true {
		n, err := fd.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				currLine += string(buf[:n])
				fmt.Printf("read: %s", currLine)
				return
			}
			fmt.Printf("error ocurred while reading the file: %v", err.Error())
		} else {
			currLine += string(buf[:n])
			parts := strings.Split(currLine, "\n")
			if len(parts) == 2 {
				fmt.Printf("read: %s\n", parts[0])
				currLine = parts[1]
			} else {
				currLine = parts[0]
			}
		}
	}
}
