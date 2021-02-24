package main

import (
	"fmt"
)

// This is a brute force approach with O(n^2) time complexity.
func twoSum(nums []int, target int) []int {
	// Iterate through the slice of numbers.
	for i, v := range nums {
		// Iterate over nums again to get ever combination of index.
		for i2, v2 := range nums {
			// Skip if the indexes are the same.
			if i == i2 {
				continue
			}
			// Return the current indexes if the corresponding values sum to target.
			if v+v2 == target {
				answer := []int{i, i2}
				return answer
			}
		}
	}
	// Return -1 if no answer is found.
	no_answer := []int{-1}
	return no_answer
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))
}
