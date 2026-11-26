package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	for i, num := range nums { // i is the index, num is the value
		fmt.Println(i, num)
	}

	m := map[string]string{"fname": "John", "lname": "Doe"}
	for k, v := range m {
		fmt.Println(k, v) // k is the key, v is the value
	}

	for i, c := range "Hello" { // i is the index, c is the rune value
		fmt.Println(i, c, string(c)) // convert rune to string
	}
}
