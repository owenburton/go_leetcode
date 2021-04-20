package main

import "fmt"

// Understand, plan, execute, review.
// Create a map with all the numbers from arr as the keys.
// Iterate through the numbers in arr.
// Double the current number and see if it's in the map.
// If so, return true. Else, false.

func checkIfExist(arr []int) bool {
	m := make(map[int]int)
	for i, v := range arr {
		if _, ok := m[v*2]; ok {
			return true
		}
		if _, ok := m[v/2]; ok && v%2==0 {
			return true
		}
		m[v] = i
	}
	return false
}

func main() {
	// arr := []int{10, 2, 5, 3}
	arr := []int{3, 1, 7, 11}

	fmt.Println(checkIfExist(arr))
}