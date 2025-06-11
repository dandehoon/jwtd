package cmd

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeToken(t *testing.T) {
	t.Parallel()
	// Valid JWT token with expiration time
	validToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiZXhwIjoxNzE2MjM5MDIyfQ.fake_signature"

	// Invalid format token
	invalidToken := "not.a.valid.token"

	// Token with invalid base64
	invalidBase64Token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.!!!invalid_base64!!!.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	tests := []struct {
		name     string
		token    string
		contains []string
	}{
		{
			name:  "valid token with expiration",
			token: validToken,
			contains: []string{
				"alg: HS256",
				"typ: JWT",
				"sub: 1234567890",
				"name: John Doe",
				"exp: 1716239022",
			},
		},
		{
			name:     "invalid token format",
			token:    invalidToken,
			contains: []string{},
		},
		{
			name:  "token with invalid base64",
			token: invalidBase64Token,
			contains: []string{
				"alg: HS256",
				"typ: JWT",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect stdout to capture output
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call function to test
			decodeToken(tt.token)

			// Reset stdout and get output
			w.Close()
			os.Stdout = oldStdout
			var buf bytes.Buffer
			io.Copy(&buf, r)
			output := buf.String()

			// Check expected output is contained
			for _, expected := range tt.contains {
				assert.Contains(t, output, expected, "Expected output to contain %q", expected)
			}
		})
	}
}
