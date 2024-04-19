package website

import (
	"html/template"
	"log"
	"path/filepath"

	"website/internal/server"
)

const port string = ":8888"

func Run() {
	// TODO: add static file

	p := filepath.Join(".", "web", "templates", "*.html")
	tmpls, err := template.ParseGlob(p)
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err)
	}

	srv := server.New(tmpls)
	srv.Routes()
	if err := srv.Start(port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
	// TODO: Graceful shutdown
}
