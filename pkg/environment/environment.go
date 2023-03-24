package environment

import "strings"

const (
	Local = "local"
	Prd   = "prd"
)

func IsValid(env string) bool {
	switch env {
	case "local", "prd":
		return true
	default:
		return false
	}
}

func FromString(env string) string {
	switch strings.ToLower(env) {
	case "local":
		return Local
	case "prd":
		return Prd
	default:
		return Local
	}
}
