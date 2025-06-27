package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain_Integration(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Setup: Save original stdin and args
	oldStdin := os.Stdin
	oldArgs := os.Args
	defer func() {
		os.Stdin = oldStdin
		os.Args = oldArgs
	}()

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	// Create mock stdin with a JWT token
	mockToken := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	mockStdin, _ := os.CreateTemp("", "mock-stdin")
	defer os.Remove(mockStdin.Name())

	_, err := mockStdin.WriteString(mockToken)
	assert.NoError(t, err)
	_, err = mockStdin.Seek(0, 0)
	assert.NoError(t, err)
	os.Stdin = mockStdin

	// Set up args for Cobra
	os.Args = []string{"jwtd"}

	// Run the main function
	main()

	// Close writer to get output
	w.Close()
	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	assert.NoError(t, err)
	output := buf.String()

	// Check expected content in output
	expectedStrings := []string{
		"alg: HS256",
		"typ: JWT",
		"sub: 1234567890",
		"name: John Doe",
		"iat: 1516239022",
	}

	for _, expected := range expectedStrings {
		assert.Contains(t, output, expected, "Expected output to contain %q", expected)
	}
}
