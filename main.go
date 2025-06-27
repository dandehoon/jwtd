package main

import (
	"log"

	"github.com/dandehoon/jwtd/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
