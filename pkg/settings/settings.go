package settings

import (
	"fmt"

	"github.com/spf13/cast"
)

func (s *settings) get(key string) any {
	value := s.source.Get(key)

	if value == nil {
		panic(fmt.Errorf("key \"%s\" is empty", key))
	}

	return value
}

func (s *settings) Get(key string) any {
	return s.source.Get(key)

}

func (s *settings) String(key string) string {
	return cast.ToString(s.get(key))
}

func (s *settings) Bool(key string) bool {
	return cast.ToBool(s.get(key))
}

func (s *settings) Int(key string) int {
	return cast.ToInt(s.get(key))
}

func (s *settings) Int64(key string) int64 {
	return cast.ToInt64(s.get(key))
}

func (s *settings) Float64(key string) float64 {
	return cast.ToFloat64(s.get(key))
}

func (s *settings) Strings(key string) []string {
	return cast.ToStringSlice(s.get(key))
}
