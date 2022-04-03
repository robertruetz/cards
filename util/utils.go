package util

import "encoding/json"

func JSONify (input interface{}) string {
	b, _ := json.Marshal(input)
	return string(b)
}