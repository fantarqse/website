package website

import (
	"log"
	"net/http"
)

const port string = ":8888"

func Run() {
	// TODO: add static file
	// TODO: init router
	log.Printf("Starting server on %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("the server failed to start: %s", err)
	}

	// TODO: Graceful shutdown
}
