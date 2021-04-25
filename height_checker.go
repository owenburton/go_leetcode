package main

import (
	"fmt"
	"sort"
)

// Understand, plan, execute, review.
// Make a copy while sorting the input array.
// Iterate through the length of the input array.
// Compare the current index value from the input
// array and sorted copy.
// If the values are not the same, increment a counter.
// Return the counter.

func heightChecker(heights []int) int {
	arr := make([]int, len(heights))
	copy(arr, heights)
	sort.Ints(arr)
	count := 0
	for i:=0; i<len(arr); i++ {
		if arr[i] != heights[i] {
			count++
		}
	}
	return count
}

func main() {
	heights := []int{1, 1, 4, 2, 1, 3}
	fmt.Println(heights)
	fmt.Println(heightChecker(heights))
}