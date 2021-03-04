package main

import "fmt"

/*
The problem and description can be found here:
https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
*/

func findDisappearedNumbers(nums []int) []int {
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

func main() {
	testSlice := []int{4, 3, 2, 7, 8, 2, 3, 1}

	fmt.Println(findDisappearedNumbers(testSlice))
}
