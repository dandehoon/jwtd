package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

// decodeToken decodes a JWT token and prints its contents
func decodeToken(token string) error {
	w := os.Stdout
	token = cleanToken(token)
	parts := strings.Split(token, ".")
	if len(parts) < 1 {
		return fmt.Errorf("invalid JWT token format")
	}

	for i, part := range parts {
		if i > 0 {
			fmt.Fprintln(w, "---")
		}

		data, err := base64.RawStdEncoding.DecodeString(part)
		if err != nil {
			log.Printf("Error decoding part %d: %v", i+1, err)
			continue
		}

		var meta map[string]any
		if err := json.Unmarshal(data, &meta); err != nil {
			continue
		}

		dict := make(map[string]string)
		keys := make(sort.StringSlice, 0, len(dict))
		for k, v := range meta {
			keys = append(keys, k)
			var ext string
			switch k {
			case "exp", "nbf", "iat":
				v, ext = getFormattedTime(v)
			}
			s := fmt.Sprintf("%s: %v %s", k, v, ext)
			dict[k] = strings.TrimSpace(s)
		}
		keys.Sort()

		for _, k := range keys {
			fmt.Fprintln(w, dict[k])
		}

	}
	return nil
}

// getFormattedTime converts a timestamp to human-readable format
func getFormattedTime(v any) (int64, string) {
	plain, ok := v.(float64)
	if !ok {
		return 0, "(invalid timestamp)"
	}
	num := int64(plain)
	ts := time.Unix(num, 0)
	return num, fmt.Sprintf("(%s)", ts.Format(time.RFC3339))
}
