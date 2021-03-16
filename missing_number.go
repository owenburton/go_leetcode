package main

import "fmt"

// Understand, plan, execute, review
// Find the sum of all nums in range len of input arr.
// Iterate through the input arr.
// Subtract each element from the sum found earlier.
// Return sum.

func missingNumber(nums []int) int {
	sum := 0
	for i := 1; i < len(nums)+1; i++ {
		sum += i
	}
	fmt.Println(sum)
	for _, v := range nums {
		sum -= v
	}
	return sum
}

func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums))
}
