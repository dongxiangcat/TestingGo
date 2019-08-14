package utils

import (
	"encoding/json"
	_ "net/http"
)

/* 所有类型转为JSON */
func EncodeJson(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}
