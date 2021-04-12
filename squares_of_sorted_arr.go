package main

import "fmt"

// Given an integer array nums sorted in non-decreasing order,
// return an array of the squares of each number sorted in non-decreasing order.
func sortedSquares(nums []int) []int {
	r := 0
	for r < len(nums) && nums[r] < 0 {
		r++
	}
	l := r - 1

	answer := make([]int, 0)

	for l >= 0 && r < len(nums) {
		lSquare := nums[l] * nums[l]
		rSquare := nums[r] * nums[r]

		if lSquare < rSquare {
			answer = append(answer, lSquare)
			l--
		} else {
			answer = append(answer, rSquare)
			r++
		}
	}

	for l >= 0 {
		answer = append(answer, nums[l]*nums[l])
		l--
	}
	for r < len(nums) {
		answer = append(answer, nums[r]*nums[r])
		r++
	}

	return answer
}

func main() {
	// nums := []int{-4, -1, 0, 3, 10}
	// nums := []int{-7, -3, 2, 3, 11}
	nums := []int{-1}
	fmt.Println(sortedSquares(nums))
}
