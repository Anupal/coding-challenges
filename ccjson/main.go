package main

import (
	"fmt"
	"github.com/anupal/coding-challenges/ccjson/json"
)

func main() {
	//testList := json.List{"1", 2}
	//
	//testList = append(testList, 3.0)
	//testList = append(testList, "anupal")
	//
	//fmt.Println(testList)
	//
	//testMap := json.Object{
	//	"hello": 12,
	//	"yono":  34,
	//	12:      3,
	//}
	//
	//testMap[6.0] = 56
	//fmt.Println(testMap)

	//sampleJSON := `{"hello": "json", "this": "is sparta"}`

	//endIndex, parsedString, err := json.ParseString(0, sampleJSON)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(endIndex, parsedString)
	//}

	sampleJSON := `{"hello1":"world1","hello2":"world2"}`
	parsedObject, err := json.ParseObject(0, sampleJSON)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(parsedObject)
	}

}
