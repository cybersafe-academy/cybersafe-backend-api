package cacheutil

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func Config(defaultExpiration, cleanupInterval time.Duration) *cache.Cache {
	return cache.New(defaultExpiration, cleanupInterval)
}
