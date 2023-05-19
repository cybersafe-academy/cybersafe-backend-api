package cacheutil

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func Config(defaultExpiration, cleanupInterval time.Duration) Cacher {
	return cache.New(defaultExpiration, cleanupInterval)
}
