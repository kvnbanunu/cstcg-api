package main

import (
	"kvnbanunu/cstcg-api/handlers"
	"kvnbanunu/cstcg-api/config"
	"log"
	"net/http"
)

func main() {
	
	cfg := config.Load()

	mux := setupRoutes()

	server := &http.Server{
		Addr: cfg.Port,
		Handler: mux,
	}

	log.Printf("Serving on port %s\n", cfg.Port)
	log.Fatal(server.ListenAndServe())
}

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HandleRoot)
	
	return mux
}
