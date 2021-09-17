package main

import (
	"fmt"
	"reflect"
)

func main() {
	var integers, evenIntegers []int
	for i := 0; i < 100; i++ {
		integers = append(integers, i)
		if i%2 == 0 {
			evenIntegers = append(evenIntegers, i)
		}
	}

	var result = hashIntersection(integers, evenIntegers)
	fmt.Println(result)
}

func hashIntersection(collectionA interface{}, collectionB interface{}) (result []interface{}) {
	// result = make([]interface{}, 0) // Is there a case where this is needed?
	hash := make(map[interface{}]bool)
	collectionValueA := reflect.ValueOf(collectionA)
	collectionValueB := reflect.ValueOf(collectionB)

	for i := 0; i < collectionValueA.Len(); i++ {
		element := collectionValueA.Index(i).Interface()
		hash[element] = true
	}

	for i := 0; i < collectionValueB.Len(); i++ {
		el := collectionValueB.Index(i).Interface()
		if _, found := hash[el]; found {
			result = append(result, el)
		}
	}

	return
}
