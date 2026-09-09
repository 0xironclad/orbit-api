package main

import (
	"log"
	"net/http"
)

type config struct {
	addr string
}

type application struct {
	config config
}

func (app *application) run() error {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
	}

	log.Printf("server has started on %s", app.config.addr)

	return server.ListenAndServe()
}
