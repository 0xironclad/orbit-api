package main

import (
	"log"
	"social/internal/env"
	"social/internal/store"
)

func main() {
	config := config{
		addr: env.GetString("ADDRESS", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: config,
		store: store,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
