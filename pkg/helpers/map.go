package helpers

import (
	"bytes"
	"encoding/json"
)

func MustMapToBytesBuffer(m map[string]any) *bytes.Buffer {
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes.NewBuffer(b)
}
