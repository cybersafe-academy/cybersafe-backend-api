package cacheutil

import (
	"cybersafe-backend-api/pkg/helpers"
	"time"
)

type CacheMock struct {
	Source helpers.M
}

func (cm *CacheMock) Get(key string) (any, bool) {
	value, found := cm.Source[key]
	return value, found
}
func (cm *CacheMock) Set(key string, value any, duration time.Duration) {
	// Duration is being ignored for testing purposes
	cm.Source[key] = value
}
func (cm *CacheMock) Delete(key string) {
	delete(cm.Source, key)
}
