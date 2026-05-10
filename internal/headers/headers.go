package headers

import (
	"bytes"
	"fmt"
)

type Headers map[string]string

func NewHeaders() Headers {
	return make(Headers)
}

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
	crlfIdx := bytes.Index(data, []byte("\r\n"))
	if crlfIdx == -1 {
		return 0, false, nil
	}
	if crlfIdx == 0 {
		return 2, true, nil
	}

	line := data[:crlfIdx]
	parts := bytes.SplitN(line, []byte(":"), 2)
	if len(parts) != 2 {
		return 0, false, fmt.Errorf("malformed header: missing colon")
	}

	key := parts[0]
	if bytes.Contains(key, []byte(" ")) {
		return 0, false, fmt.Errorf("malformed header: space before colon in %q", key)
	}

	h[string(bytes.TrimSpace(key))] = string(bytes.TrimSpace(parts[1]))
	return crlfIdx + 2, false, nil
}
