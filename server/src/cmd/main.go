// main.go
package main

import (
	"log"
	"net/http"

	"github.com/Deathfireofdoom/terraxcel/server/src/api"
)

func main() {
	// Initialize router
	r := api.NewRouter()

	// Run server
	server := &http.Server{Addr: ":8080", Handler: r}

	log.Println("Starting server on port 8080...")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed", err)
	}

}
