package cacheutil

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var c *cache.Cache

func Config(defaultExpiration, cleanupInterval time.Duration) {
	c = cache.New(defaultExpiration, cleanupInterval)
}

func GetCache() *cache.Cache {
	return c
}
