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

		if i < len(parts)-2 {
			fmt.Fprintln(w, "---")
		}
	}
}

func getFormattedTime(v any) (int64, string) {
	plain, ok := v.(float64)
	if !ok {
		return 0, "(invalid timestamp)"
	}
	num := int64(plain)
	ts := time.Unix(num, 0)
	return num, ts.Format(time.RFC3339)
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
