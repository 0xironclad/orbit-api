package main

import (
	"log"
	"social/internal/database"
	"social/internal/env"
	"social/internal/store"
)

func main() {
	config := config{
		addr: env.GetString("ADDRESS", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDRESS", "postgres://user:password@localhost:5432/dbname"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 20),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 20),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "10min"),
		},
	}

	
	db, err := database.New(config.db.addr, config.db.maxOpenConns, config.db.maxIdleConns, config.db.maxIdleTime)
	if err != nil {
		log.Fatal(err)
	}
	
	store := store.NewStorage(db)

	app := &application{
		config: config,
		store: store,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
