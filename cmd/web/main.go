package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hidayattaufiqur/go-course/pkg/config"
	"github.com/hidayattaufiqur/go-course/pkg/handlers"
	"github.com/hidayattaufiqur/go-course/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
