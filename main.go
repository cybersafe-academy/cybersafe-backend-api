package main

import (
	_ "cybersafe-backend-api/docs"

	"cybersafe-backend-api/pkg/components"
	"cybersafe-backend-api/pkg/server"
	"fmt"
	"log"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	port := "8080"
	log.Printf("Starting up on http://localhost:%s", port)

	c := components.Config()

	server.Config(c)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), c.Router))
}
