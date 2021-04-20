package main

import "fmt"

// Understand, plan, execute, review.
// Create a map with all the numbers from arr as the keys.
// Iterate through the numbers in arr.
// Double the current number and see if it's in the map.
// If so, return true. Else, false.

func checkIfExist(arr []int) bool {
	var cache = map[int]int{}
	for i, v := range arr {
		cache[v] = i
	}
	for i, v := range arr {
		if idx, ok := cache[v*2]; ok {
			if i!=idx {
				return true
			} 
		}
	}
	return false
}

func main() {
	// arr := []int{10, 2, 5, 3}
	arr := []int{3, 1, 7, 11}

	fmt.Println(checkIfExist(arr))
}