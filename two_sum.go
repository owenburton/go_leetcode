package main

import (
	"fmt"
)

/*
The problem & description for two_sum can be found here:
https://leetcode.com/problems/two-sum/
*/

// This is a brute force approach with O(n^2) time complexity.
func twoSum1(nums []int, target int) []int {
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
	noAnswer := []int{-1}
	return noAnswer
}

// This is usually an O(n) solution using a hash table.
func twoSum2(nums []int, target int) []int {
	// Store numbers and their indexes in hash table.
	numsIndexLookup := make(map[int][]int)
	for i, v := range nums {
		// Append i if v is already in map.
		if _, ok := numsIndexLookup[v]; ok {
			prev := numsIndexLookup[v]
			numsIndexLookup[v] = append(prev, i)
		} else {
			indexSlice := []int{i}
			numsIndexLookup[v] = indexSlice
		}
	}

	// Iterate through nums to check for a solution.
	for i, v := range nums {
		otherVal := target - v

		if i2, ok := numsIndexLookup[otherVal]; ok {
			answer := []int{i, i2[0]}

			// Get next index in slice if current nums value is the same.
			if v == otherVal {
				// If it found the current number, but there's no duplcate,
				// then continue to next number in nums.
				if len(i2) == 1 {
					continue
				}
				answer[1] = i2[1]
			}

			return answer
		}
	}
	// Return -1 if no answer is found.
	noAnswer := []int{-1}
	return noAnswer
}

// This is a one-pass O(n) solution using a hash table.
func twoSum3(nums []int, target int) []int {
	// Create map to store nums values & indexes.
	indexLookup := make(map[int]int)

	for i, v := range nums {
		// If the target less curret value is in the cache map,
		// return that value's index and the current one.
		if i2, ok := indexLookup[target-v]; ok {
			answer := []int{i, i2}
			return answer
		}

		// Add current value and index to cache map.
		indexLookup[v] = i
	}

	// Return -1 if no answer was found.
	noAnswer := []int{-1}
	return noAnswer
}

func main() {
	// nums := []int{2, 7, 11, 15}
	// target := 9

	// fmt.Println(twoSum(nums, target))

	nums := []int{3, 2, 4}
	target := 6

	fmt.Println(twoSum3(nums, target))
}
