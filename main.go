package main

import (
	"log"

	"github.com/danztran/jwtd/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
