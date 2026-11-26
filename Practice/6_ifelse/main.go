package main

import "fmt"

func main() {
	age := 18
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are not an adult.")
	}

	if rent := 100; rent > 1000 {
		fmt.Println("Rent is too high.")
	} else if rent < 500 {
		fmt.Println("Rent is affordable.")
	} else {
		fmt.Println("Rent is moderate.")
	}
}
