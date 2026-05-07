package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	inputFilePath := "messages.txt"
	f, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("could not open %s: %s\n", inputFilePath, err)
	}
	defer f.Close()

	fmt.Printf("Reading data from %s\n", inputFilePath)
	fmt.Println("=====================================")

	currentLineContents := ""
	buffer := make([]byte, 8)
	for {
		n, err := f.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				currentLineContents += string(buffer[:n])
				if currentLineContents != "" {
					fmt.Printf("read: %s\n", currentLineContents)
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
			fmt.Printf("read: %s%s\n", currentLineContents, parts[i])
			currentLineContents = ""
		}
		currentLineContents += parts[len(parts)-1]
	}
}
