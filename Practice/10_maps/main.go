package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]string)
	m["key1"] = "value1"
	m["key2"] = "value2"
	m["key3"] = "value3"
	m["key4"] = "value4"
	fmt.Println(m)
	fmt.Println(m["key1"])
	fmt.Println(m["key2"])
	fmt.Println(m["key5"]) // returns zero value for string type, which is an empty string

	delete(m, "key1") // deletes the key-value
	fmt.Println(m)
	clear(m) // clears the map
	fmt.Println(m)

	n := map[string]int{"key1": 1, "key2": 1, "key3": 3, "key4": 4}
	fmt.Println(n)

	value, ok := n["key5"] // checks if the key exists in the map
	if ok {
		fmt.Println("key1 exists in the map with value:", value)
	} else {
		fmt.Println("key1 does not exist in the map")
	}

	fmt.Println(maps.Equal(m, m))
}
