package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cast"
)

func main() {
	token := getToken()
	if token != "" {
		decodeToken(token)
	}
}

func decodeToken(token string) {
	w := os.Stdout
	token = cleanToken(token)

	parts := strings.Split(token, ".")

	for i, part := range parts {
		data, err := base64.RawStdEncoding.DecodeString(part)
		if err != nil {
			continue
		}

		var meta map[string]interface{}
		if err := json.Unmarshal(data, &meta); err != nil {
			continue
		}

		dict := make(map[string]string)
		keys := make(sort.StringSlice, 0, len(dict))
		for k, v := range meta {
			keys = append(keys, k)
			var s string
			switch k {
			case "exp", "nbf", "iat":
				num := cast.ToInt64(v)
				ts := time.Unix(num, 0)
				s = fmt.Sprintf("%s: %d (%s)", k, num, ts.Format(time.RFC3339))
			default:
				s = fmt.Sprintf("%s: %v", k, v)
			}
			dict[k] = s
		}
		keys.Sort()

		for _, k := range keys {
			fmt.Fprintln(w, dict[k])
		}

		if i < len(parts)-2 {
			fmt.Fprintln(w, "---")
		}
	}
}

func getToken() string {
	d := os.Stdin
	defer d.Close()
	s, err := io.ReadAll(d)
	if err != nil {
		panic(err)
	}
	return string(s)
}

func cleanToken(token string) string {
	token = strings.TrimSpace(token)
	parts := strings.SplitN(token, " ", 2)
	token = parts[len(parts)-1]
	token = strings.TrimSpace(token)
	return token
}
