package main

import "fmt"

// Given an integer array nums sorted in non-decreasing order,
// return an array of the squares of each number sorted in non-decreasing order.
func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	l := 0
	r := n - 1

	for i := n - 1; i >= 0; i-- {
		lSqr := nums[l] * nums[l]
		rSqr := nums[r] * nums[r]

		if lSqr > rSqr {
			ans[i] = lSqr
			l++
		} else {
			ans[i] = rSqr
			r--
		}
	}

	return ans
}

func main() {
	// nums := []int{-4, -1, 0, 3, 10}
	nums := []int{-7, -3, 2, 3, 11}
	// nums := []int{-1}
	fmt.Println(sortedSquares(nums))
}
