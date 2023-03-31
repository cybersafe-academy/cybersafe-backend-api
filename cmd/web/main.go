package main

import (
	_ "cybersafe-backend-api/docs"
	"cybersafe-backend-api/pkg/components"
	"cybersafe-backend-api/pkg/server"
	"fmt"
	"log"
	"net/http"
)

//	@title						CyberSafe Academy APIs
//	@Description				REST Api for all the system services
//	@version					0.001Beta
//	@BasePath					/api/v1
//	@Accept						json
//	@Produce					json
//	@query.collection.format	multi

//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token. e.g: Bearer eyJhbGciO....

// @securityDefinitions.apikey	Language
// @in							header
// @name						Accept-Language
func main() {
	port := "8080"
	log.Printf("Starting up on http://localhost:%s", port)

	c := components.Config()

	server.Config(c)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), c.Router))
}
