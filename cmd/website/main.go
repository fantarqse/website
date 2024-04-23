package main

import (
	"log"

	website "website/cmd"
)

func main() {
	if err := website.Run(); err != nil {
		log.Fatalf("failed to start server: %s", err.Error())
	}
}
