package headers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequestHeadersParse(t *testing.T) {
	// Test: Valid single header
	headers := NewHeaders()
	data := []byte("Host: localhost:42069\r\n\r\n")
	n, done, err := headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Equal(t, "localhost:42069", headers["Host"])
	assert.Equal(t, 23, n)
	assert.False(t, done)

	// Test: Invalid spacing header
	headers = NewHeaders()
	data = []byte("       Host: localhost:42069\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	t.Run("Valid single header", func(t *testing.T) {
		headers := NewHeaders()
		n, done, err := headers.Parse([]byte("Host: localhost\r\n"))
		require.NoError(t, err)
		assert.Equal(t, 17, n)
		assert.False(t, done)
		assert.Equal(t, "localhost", headers["Host"])
	})

	t.Run("Valid single header with extra whitespace", func(t *testing.T) {
		headers := NewHeaders()
		n, done, err := headers.Parse([]byte("Host:   localhost  \r\n"))
		require.NoError(t, err)
		assert.Equal(t, 21, n)
		assert.False(t, done)
		assert.Equal(t, "localhost", headers["Host"])
	})

	t.Run("Valid 2 headers with existing headers", func(t *testing.T) {
		headers := NewHeaders()
		headers["Host"] = "localhost"
		n, done, err := headers.Parse([]byte("Content-Type: application/json\r\n"))
		require.NoError(t, err)
		assert.Equal(t, 32, n)
		assert.False(t, done)
		assert.Equal(t, "localhost", headers["Host"])
		assert.Equal(t, "application/json", headers["Content-Type"])
	})

	t.Run("Valid done", func(t *testing.T) {
		headers := NewHeaders()
		n, done, err := headers.Parse([]byte("\r\n"))
		require.NoError(t, err)
		assert.Equal(t, 2, n)
		assert.True(t, done)
	})

	t.Run("Invalid spacing header", func(t *testing.T) {
		headers := NewHeaders()
		n, done, err := headers.Parse([]byte("Host : localhost\r\n"))
		require.Error(t, err)
		assert.Equal(t, 0, n)
		assert.False(t, done)
	})

}
