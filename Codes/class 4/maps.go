package main

import "fmt"

func main() {
	// map is a collection of data
	// where we have a key-value pair
	// wherever you need to do searching on a collection of data, you should use map
	// map[data_type_of_key]data_type_of_value

	myMap := make(map[string]int)

	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	myMap["four"] = 4
	myMap["five"] = 5
	myMap["six"] = 6
	myMap["seven"] = 7
	myMap["eight"] = 8
	myMap["a"] = 8
	myMap["z"] = 8

	fmt.Println("Printed by print function", myMap)
	fmt.Println("Printing myMap by loop")
	for key, val := range myMap {
		fmt.Println(key, " : ", val)
	}

	anotherMap := make(map[bool]string)
	anotherMap[true] = "string_for_true"
	fmt.Println(anotherMap[true])

	value, isKeyPresent := myMap["one"]
	if isKeyPresent {
		fmt.Println("key is present and the value to it is", value)
	} else {
		fmt.Println("key not present")
	}

}
