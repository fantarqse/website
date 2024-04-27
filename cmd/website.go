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

	if err := server.New(tmpls).Routes().Start(port); err != nil {
		return err
	}

	return nil
}
