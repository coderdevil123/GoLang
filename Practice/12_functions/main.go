package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func getLanguages() (string, string, string) {
	return "golang", "python", "java"
}

func processIt(fn func(int, int) int, x int, y int) int {
	return fn(x, y)
}

func main() {
	fmt.Println("Hello, World!")
	result := add(5, 3)
	fmt.Println("The result of addition is:", result)
	// lang1, lang2, lang3 := getLanguages()
	// fmt.Println("Languages:", lang1, lang2, lang3)
	result2 := processIt(add, 5, 3)
	fmt.Println("The result of processing is:", result2)
	lang1, lang2, _ := getLanguages() // ignoring the third return value
	fmt.Println("Languages:", lang1, lang2)
}
