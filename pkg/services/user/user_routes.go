package user

import (
	"cybersafe-backend-api/pkg/components"
	"cybersafe-backend-api/pkg/db"
)

func GetAll(c *components.HTTPComponents) {

	_, _ = db.GetDatabaseConnection()

	c.HttpResponse.Write([]byte(c.Components.Environment))
}
