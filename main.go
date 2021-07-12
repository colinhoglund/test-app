package main

import (
	"log"
	"net/http"
	"os"

	"github.com/colinhoglund/test-app/internal/server"
)

func main() {
	addr := os.Getenv("APP_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	log.Printf("Starting server on %s\n", addr)
	log.Fatalln(http.ListenAndServe(addr, server.New()))
}
