package internationalization

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"golang.org/x/text/language"
)

var Bundle *i18n.Bundle

func Config() *i18n.Bundle {

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("pkg/internationalization/en.toml")
	bundle.MustLoadMessageFile("pkg/internationalization/pt.toml")

	Bundle = bundle

	return bundle
}
