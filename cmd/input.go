package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// getToken retrieves JWT token from stdin or command line arguments
func getToken(args []string) (string, error) {
	// First try to read from stdin
	stat, err := os.Stdin.Stat()
	if err == nil && (stat.Mode()&os.ModeCharDevice) == 0 {
		// Data is available on stdin
		s, err := io.ReadAll(os.Stdin)
		if err == nil && len(strings.TrimSpace(string(s))) > 0 {
			return string(s), nil
		}
	}

	// If stdin is empty or unavailable, try command line arguments
	if len(args) > 0 {
		token := strings.TrimSpace(args[0])
		if token != "" {
			return token, nil
		}
	}

	return "", fmt.Errorf("no token provided")
}

// cleanToken removes prefixes and extra whitespace from JWT token
func cleanToken(token string) string {
	token = strings.TrimSpace(token)
	parts := strings.SplitN(token, " ", 2)
	token = parts[len(parts)-1]
	token = strings.TrimSpace(token)
	return token
}
