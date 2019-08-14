package main

import (
	"encoding/json"
	"fmt"
)

type Test struct {
	Message string
}

func init() {
	b, _ := json.Marshal(Test{"abc"})
	fmt.Println(string(b))

}
