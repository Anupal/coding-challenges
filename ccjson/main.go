package main

import (
	"fmt"
	"github.com/anupal/coding-challenges/ccjson/json"
)

func main() {
	//i, obj, err := json.ParseObject(0, `{"a":{"b":"c"},"d":"e"}`)
	//i, obj, err := json.ParseObject(0, `{"a":["b","c"],"d":"e"}`)
	//fmt.Println(i, obj, err)

	i, arr, err := json.ParseArray(0, `["b","c",["b","c"]]`)
	fmt.Println(i, arr, err)
}
