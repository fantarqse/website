package website

import (
	"html/template"
	"path/filepath"

	"website/internal/server"
)

const port string = ":8888"

func Run() error {
	p := filepath.Join(".", "web", "templates", "*.html")
	tmpls, err := template.ParseGlob(p)
	if err != nil {
		return err
	}

	srv := server.New(tmpls)
	srv.Routes()
	if err := srv.Start(port); err != nil {
		return err
	}

	return nil
}
