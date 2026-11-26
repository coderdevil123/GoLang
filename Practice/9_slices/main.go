package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums []int // uninitialized slice is nil
	fmt.Println(nums)
	fmt.Println(len(nums))

	var nums2 = make([]int, 2) // initialized slice with length 2
	fmt.Println(nums2)
	fmt.Println(len(nums2))
	fmt.Println(cap(nums2)) // capacity of the slice is the maximum number of elements it can hold before needing to allocate more memory

	var nums3 = make([]int, 2, 5) // initialized slice with length 2 and capacity 5
	fmt.Println(nums3)
	fmt.Println(len(nums3))
	fmt.Println(cap(nums3))
	nums3 = append(nums3, 1, 2, 3) // appending elements to the slice
	fmt.Println(nums3)
	fmt.Println(len(nums3))
	fmt.Println(cap(nums3))

	nums4 := []int{}
	nums4 = append(nums4, 1, 2, 3, 4, 5)
	fmt.Println(nums4)
	fmt.Println(len(nums4))
	fmt.Println(cap(nums4))

	var nums5 = make([]int, 0, 5)
	nums5 = append(nums5, 1)
	var nums6 = make([]int, len(nums5))
	copy(nums6, nums5)
	fmt.Println(nums5, nums6)

	fmt.Println(nums4[0:3])

	fmt.Println(slices.Equal(nums5, nums6))

	var nums7 = [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}
	fmt.Println(nums7)
}
