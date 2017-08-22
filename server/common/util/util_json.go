package util

import "encoding/json"

func JsonToString(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
