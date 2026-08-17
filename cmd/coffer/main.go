package main

import (
	"log"

	"github.com/joshcwainwright/coffer/internal/config"
	"github.com/joshcwainwright/coffer/internal/store"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failure loading config: %v", err)
	}

	st, err := store.Open(cfg.DBPath)
	if err != nil {
		log.Fatalf("failure openening store: %v", err)
	}
	defer func() { _ = st.Close() }()

	log.Printf("database ready at %s", cfg.DBPath)
}
