package cacheutil

import "time"

type Cacher interface {
	Get(string) (any, bool)
	Set(string, any, time.Duration)
}
