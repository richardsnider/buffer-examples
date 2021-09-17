package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type object map[string]interface{}

type structure struct {
	prop1 interface{}
}

func customCopyMap(original object) (copied object) {
	copied = make(object)
	for key, value := range original {
		assertedValue, ok := value.(object)
		if ok {
			copied[key] = customCopyMap(assertedValue)
		} else {
			copied[key] = value
		}
	}

	return
}

func alterMap(param object) {
	var paramCopy = make(object)
	copier.CopyWithOption(&paramCopy, &param, copier.Option{DeepCopy: true})
	paramCopy["prop1"] = 0
	var innerProp = paramCopy["innerProp"].(object)
	innerProp["prop2"] = "bar"
	fmt.Println(paramCopy)
}

func alterStructure(param structure) {
	param.prop1 = 0
}

func main() {
	var exampleObject = object{
		"prop1": 42,
		"innerProp": object{
			"prop1": 9001,
			"prop2": "foo",
		},
	}
	alterMap(exampleObject)
	fmt.Println(exampleObject)

	var exampleStructure = structure{prop1: 42}
	alterStructure(exampleStructure)
	fmt.Println(exampleStructure)
}
