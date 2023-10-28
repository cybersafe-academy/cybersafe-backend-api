package settings

import (
	"github.com/spf13/viper"
)

type Settings interface {
	Get(string) any
	String(string) string
	Float64(string) float64
	Int(string) int
	Int64(string) int64
	Bool(string) bool
	Strings(string) []string
	StrWDefault(string, string) string
}

type settings struct {
	source *viper.Viper
}
