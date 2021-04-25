package main

import "fmt"

// Understand, Plan, Execute, Review
// Bubble the zeroes to the end of the array.
// Start a while loop.
// Iterate through the array.
// If the current element is a zero and the next isn't, switch them.
// Have a boolean value called sorted.
// If the above condition was met, change sorted bool to true.
// Stop the while loop when sorted bool is true.

// func moveZeroes(nums []int) {
// 	sorted := false
// 	for sorted != true {
// 		sorted = true
// 		for i := 1; i < len(nums); i++ {
// 			if nums[i-1] == 0 && nums[i] != 0 {
// 				nums[i-1] = nums[i]
// 				nums[i] = 0
// 				sorted = false
// 			}
// 		}
// 	}
// }

// Understand, plan, execute, review.
// Make a write pointer.
// Iterate over the array.
// If the current value is not a zero,
// change the value at index write pointer to the current value,
// then change the current value to zero, and
// increment the write pointer.

func moveZeroes( nums []int) {
	idx := 0
	for i, v := range nums {
		if v != 0 {
			nums[i] = 0
			nums[idx] = v
			idx ++
		}
	}
}

func main() {
	// nums := []int{0, 1, 0, 3, 12}
	nums := []int{1}
	// nums := []int{1, 0}

	fmt.Println(nums)
	moveZeroes(nums)
	fmt.Println(nums)
}
