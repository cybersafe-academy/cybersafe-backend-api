package main

import (
	"cybersafe-backend-api/pkg/components"
	"cybersafe-backend-api/pkg/server"
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

//	@security					ApiKeyAuth
//	@securityDefinitions.basic	BasicAuth

//	@securitydefinitions.oauth2.application	OAuth2Application
//	@tokenUrl								https://example.com/oauth/token

func main() {

	c := components.Config()

	server.Config(c)

	log.Printf("Starting up on %s:%s", c.Settings.String("docs.host"), c.Settings.String("docs.port"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", c.Settings.String("docs.port")), c.Router))
}
