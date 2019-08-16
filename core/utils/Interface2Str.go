package utils

import (
	"fmt"
	"strconv"
)

func ToString(v interface{}) string {
	fmt.Println("进入转换")
	var res string
	switch v.(type) {
	case nil:
		res = ""
	case int:
		res = strconv.Itoa(v.(int))
	case int64:
		res = strconv.Itoa(int(v.(int64)))
	case int32:
		res = strconv.Itoa(int(v.(int32)))
	case int16:
		res = strconv.Itoa(int(v.(int16)))
	case int8:
		res = strconv.Itoa(int(v.(int8)))
	case float32:
		res = strconv.Itoa(int(v.(float32)))
	case float64:
		res = strconv.Itoa(int(v.(float64)))
	case string:
		res = v.(string)
	default:
		res = ""
	}
	return res
}
