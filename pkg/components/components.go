package components

import (
	"cybersafe-backend-api/pkg/db"
	"cybersafe-backend-api/pkg/environment"
	"cybersafe-backend-api/pkg/logger"
	"cybersafe-backend-api/pkg/settings"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
)

type Components struct {
	Environment string
	Router      *chi.Mux
	Logger      *zerolog.Logger
	Settings    settings.Settings
}

type HTTPComponents struct {
	HttpResponse http.ResponseWriter
	HttpRequest  *http.Request
	Components   *Components
}

func Config() *Components {
	var applications []string

	env := os.Getenv("ENV")

	applications = append(applications, "configs/application.yml")

	if environment.IsValid(env) {
		applications = append(applications, fmt.Sprintf("configs/application_%s.yml", environment.FromString(env)))
	}

	config := settings.Config("", applications)
	log := logger.Config("/", config.String("application.name"), "v1", (env == environment.Prd))

	db.CreateDBConnection(config)

	err := db.AutoMigrateDB()

	if err != nil {
		log.Info().Err(err).Msg("Error occurred while trying to run migrations...")
		os.Exit(-1)
	}

	return &Components{
		Settings:    config,
		Environment: environment.FromString(env),
		Logger:      log,
	}
}

func HttpComponents(writer http.ResponseWriter, request *http.Request, c *Components) *HTTPComponents {
	httpComp := HTTPComponents{
		HttpRequest:  request,
		HttpResponse: writer,
		Components:   c,
	}
	return &httpComp
}
