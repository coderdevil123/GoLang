package main

import "fmt"

func main() {
	var nums [4]int
	fmt.Println(len(nums))

	nums[0] = 1
	nums[1] = 2

	fmt.Println(nums[0])
	fmt.Println(nums[1])

	fmt.Println(nums)

	dig := [3]int{1, 2, 3}
	fmt.Println(dig)

	price := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(price)
}
