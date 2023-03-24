package main

import (
	"cybersafe-backend-api/pkg/components"
	"cybersafe-backend-api/pkg/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := "8080"
	log.Printf("Starting up on http://localhost:%s", port)

	c := components.Config()

	server.Config(c)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), c.Router))
}
