package settings

import (
	"fmt"

	"github.com/spf13/cast"
)

type SettingsMock struct {
	Source map[string]any
}

func (sm *SettingsMock) get(key string) any {
	value := sm.Source[key]

	if value == nil {
		panic(fmt.Errorf("key \"%s\" is empty", key))
	}

	return value
}

func (sm *SettingsMock) getWithDefault(key string, dflt any) any {
	value := sm.Get(key)

	if value == nil {
		return dflt
	}

	return value
}

func (sm *SettingsMock) Get(key string) any {
	return sm.get(key)
}

func (sm *SettingsMock) String(key string) string {
	return cast.ToString(sm.get(key))
}

func (sm *SettingsMock) StrWDefault(key string, dflt string) string {
	return cast.ToString(sm.getWithDefault(key, dflt))
}

func (sm *SettingsMock) Bool(key string) bool {
	return cast.ToBool(sm.get(key))
}

func (sm *SettingsMock) Int(key string) int {
	return cast.ToInt(sm.get(key))
}

func (sm *SettingsMock) Int64(key string) int64 {
	return cast.ToInt64(sm.get(key))
}

func (sm *SettingsMock) Float64(key string) float64 {
	return cast.ToFloat64(sm.get(key))
}

func (sm *SettingsMock) Strings(key string) []string {
	return cast.ToStringSlice(sm.get(key))
}
