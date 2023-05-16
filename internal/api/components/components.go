package components

import (
	"cybersafe-backend-api/docs"
	"cybersafe-backend-api/internal/services"
	"cybersafe-backend-api/internal/services/courses"
	"cybersafe-backend-api/internal/services/users"
	"cybersafe-backend-api/pkg/cacheutil"
	"cybersafe-backend-api/pkg/db"
	"cybersafe-backend-api/pkg/environment"
	"cybersafe-backend-api/pkg/logger"
	"cybersafe-backend-api/pkg/mail"
	"cybersafe-backend-api/pkg/settings"
	"cybersafe-backend-api/pkg/validation"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
)

type Components struct {
	Environment string
	Router      *chi.Mux
	Logger      *zerolog.Logger
	Settings    settings.Settings
	Resources   services.Resources
	Cache       cacheutil.Cacher
	Mail        mail.Mailer
}

type HTTPComponents struct {
	HttpResponse http.ResponseWriter
	HttpRequest  *http.Request
	Components   *Components
}

func Config() *Components {

	// ENV
	env := os.Getenv("ENV")

	// Settings
	applications := []string{"configs/application.yml"}
	if environment.IsValid(env) {
		applications = append(applications, fmt.Sprintf("configs/application_%s.yml", environment.FromString(env)))
	}

	config := settings.Config("", applications)

	//Logger
	log := logger.Config("/", config.String("application.name"), "v1", (env == environment.Prd))

	// Swagger Docs
	docs.SwaggerInfo.Host = fmt.Sprintf(
		"%s:%s",
		config.StrWDefault("docs.host", "localhost"),
		config.StrWDefault("docs.port", "8080"),
	)

	docs.SwaggerInfo.BasePath = config.StrWDefault("docs.basePath", "/api")

	//Cache
	cache := cacheutil.Config(1*time.Hour, 30*time.Minute)

	// Validation
	validation.Config()

	//Mail
	mailer := mail.Config(config)

	// Database Connection
	dbConn := db.CreateDBConnection(config)

	//Migrations
	err := db.AutoMigrateDB()
	if err != nil {
		panic("Error occurred while trying to run migrations...")

	}

	return &Components{
		Settings:    config,
		Environment: environment.FromString(env),
		Logger:      log,
		Cache:       cache,
		Mail:        mailer,
		Resources: services.Resources{
			Users:   users.Config(dbConn),
			Courses: courses.Config(dbConn),
		},
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
