package request

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	req := &Request{}
	// read request line
	buf := bufio.NewReader(reader)
	line, err := buf.ReadString('\n')
	if err != nil {
		return nil, err
	}
	line = strings.TrimRight(line, "\r\n")
	ReqLine, err := parseRequestLine(line)
	if err != nil {
		return nil, err
	}
	req.RequestLine = *ReqLine
	return req, nil
}

func parseRequestLine(line string) (*RequestLine, error) {
	rl := &RequestLine{}
	// request line: method sp /path sp http-version
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("malformed request line")
	}
	rl.Method = parts[0]
	for _, ch := range rl.Method {
		if !unicode.IsUpper(ch) || !unicode.IsLetter(ch) {
			return nil, fmt.Errorf("malformed request line: method does not contain only capital alpahabetic characters: %s", rl.Method)
		}
	}
	rl.RequestTarget = parts[1]
	HttpVersion := parts[2]
	ps := strings.Split(HttpVersion, "/")
	if len(ps) == 2 {
		version := ps[1]
		if version != "1.1" {
			return nil, fmt.Errorf("malformed request line: invalid version: %s", HttpVersion)
		}
		rl.HttpVersion = version
	} else {
		return nil, fmt.Errorf("malformed request line: invalid version: %s", HttpVersion)
	}

	return rl, nil
}
