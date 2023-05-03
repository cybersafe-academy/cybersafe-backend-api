package components

import (
	"cybersafe-backend-api/docs"
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
	"gorm.io/gorm"
)

// type Services struct {
// 	Users user.UserManager
// }

type Components struct {
	Environment string
	Router      *chi.Mux
	Logger      *zerolog.Logger
	Settings    settings.Settings
	Storer      db.Storer
	DB          *gorm.DB
	Cache       *cacheutil.Cacher
	Mail        *mail.Mailer
}

type HTTPComponents struct {
	HttpResponse http.ResponseWriter
	HttpRequest  *http.Request
	Components   *Components
}

func Config() *Components {

	env := os.Getenv("ENV")

	applications := []string{"configs/application.yml"}

	if environment.IsValid(env) {
		applications = append(applications, fmt.Sprintf("configs/application_%s.yml", environment.FromString(env)))
	}

	config := settings.Config("", applications)

	settings.ExportedSettings = config

	log := logger.Config("/", config.String("application.name"), "v1", (env == environment.Prd))

	docs.SwaggerInfo.Host = fmt.Sprintf(
		"%s:%s",
		config.StrWDefault("docs.host", "localhost"),
		config.StrWDefault("docs.port", "8080"),
	)

	docs.SwaggerInfo.BasePath = config.StrWDefault("docs.basePath", "/api")

	cache := cacheutil.Config(1*time.Hour, 30*time.Minute)

	validation.Config()
	dbConn := db.CreateDBConnection(config)
	err := db.AutoMigrateDB()

	mailer := mail.Config(config)

	if err != nil {
		panic("Error occurred while trying to run migrations...")
	}

	return &Components{
		Settings:    config,
		Environment: environment.FromString(env),
		Logger:      log,
		DB:          dbConn,
		Cache:       &cache,
		Mail:        &mailer,
		Storer:      &db.DBStorer{Conn: dbConn},
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

func GetSettings() {

}
