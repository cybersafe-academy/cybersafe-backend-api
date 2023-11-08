package main

import (
	"cybersafe-backend-api/internal/api/components"
	"cybersafe-backend-api/internal/api/server"
	"fmt"
	"log"
	"net/http"
)

//	@title			CyberSafe Academy API
//	@version		1.0
//	@description	This REST API contains all services for the CyberSafe plataform.
//	@termsOfService	http://cybersafe.academy.com/support/terms

//	@contact.name	CyberSafe support team
//	@contact.url	http://cybersafe.academy.com/support/contact
//	@contact.email	support@cybersafe.com

//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit/

//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@description				Insert the token without "Bearer" prefix.
func main() {

	c := components.Config()

	server.Config(c)

	log.Printf("Starting up on %s:%s", c.Settings.String("docs.host"), c.Settings.String("docs.port"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", c.Settings.String("docs.port")), c.Router))
}
