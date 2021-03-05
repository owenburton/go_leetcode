package main

import "fmt"

/*
The problem and description can be found here:
https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
*/

func findDisappearedNumbers1(nums []int) []int {
	// Understand, plan, execute, review
	// Make a hash map with keys as numbers in nums
	// Iterate through a range of numbers from 1 to len of nums
	// if the current number is not in the cache, add it to answer array
	// return the answer array

	cache := make(map[int]int)
	for _, num := range nums {
		cache[num] = 0
	}

	var answer []int
	for num := 1; num <= len(nums); num++ {
		if _, ok := cache[num]; ok {
			continue
		} else {
			answer = append(answer, num)
		}
	}

	return answer

}

// This solution doesn't use extra space (other than the answer Slice).
func findDisappearedNumbers2(nums []int) []int {
	// Understand, plan, execute, review
	// Iterate through the numbers in nums.
	// Use the current number minus one as an index.
	// Change the number at that index to negative.
	// Initialize an empty slice to return the answer.
	// Iterate through nums again and add all positive numbers to answer.

	for _, num := range nums {
		if num < 0 {
			num *= -1
		}

		if nums[num-1] > 0 {
			nums[num-1] *= -1
		}
	}

	var answer []int
	for i, num := range nums {
		if num > 0 {
			answer = append(answer, i+1)
		}
	}

	return answer
}

func main() {
	testSlice := []int{4, 3, 2, 7, 8, 2, 3, 1}

	fmt.Println(findDisappearedNumbers2(testSlice))
}
