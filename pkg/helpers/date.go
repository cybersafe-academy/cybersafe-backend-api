package helpers

import "time"

func DefaultTimeZone() time.Time {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return time.Now()
	}
	return time.Now().In(loc)
}

func DefaultDateFormat() string {
	return "2006-01-02"
}
