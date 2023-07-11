package utils

import (
	"encoding/json"
)

func StructToJsonString(s interface{}) string {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}
