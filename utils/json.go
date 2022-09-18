package utils

import (
	"encoding/json"
)

// ToJSON bean->json
func ToJSON(bean interface{}) (j string) {
	var (
		bytes []byte
	)
	bytes, _ = json.Marshal(bean)

	j = string(bytes)

	return
}
