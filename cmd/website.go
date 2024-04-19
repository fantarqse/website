package website

import (
	"log"

	"website/internal/server"
)

const port string = ":8888"

func Run() {
	// TODO: add static file
	srv := server.New()
	srv.Routes()

	log.Printf("Starting server on %s\n", port)
	if err := srv.Start(port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

	// TODO: Graceful shutdown
}
