package helpers

import "regexp"

func RemoveAllSpecialChars(str string) string {
	return regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(str, "")
}
