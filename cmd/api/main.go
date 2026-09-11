package main

import (
	"log"
	"social/internal/env"
)

func main() {
	config := config{
		addr: env.GetString("ADDRESS", ":8080"),
	}

	app := &application{
		config: config,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
