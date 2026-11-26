package main

import "fmt"

func main() {
	const name string = "John"
	const age = 30
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	const (
		port = 1000
		host = "localhost"
	)
	fmt.Println(port, host)
}
