package main

import (
	"log"

	"github.com/joshcwainwright/net-worth/api/internal/config"
	"github.com/joshcwainwright/net-worth/api/internal/store"
)

func main() {
	cfg := config.Load()

	st, err := store.Open(cfg.DBPath)
	if err != nil {
		log.Fatalf("open store: %v", err)
	}
	defer func() { _ = st.Close() }()

	log.Printf("database ready at %s", cfg.DBPath)
}
