package main

import "fmt"

// Understand, plan, execute, review.
// Create a write pointer.
// Iterate over the idxs and values of the array.
// If the current value is even,
// Switch it with the value at the write pointer.

func sortArrayByParity(A []int) []int {
	idx := 0
	for i, v := range A {
		if v % 2 == 0 {
			A[i] = A[idx]
			A[idx] = v
			idx++
		}
	}
	return A
}

func main() {
	A := []int{3, 1, 2, 4}
	fmt.Println(A)
	fmt.Println(sortArrayByParity(A))
}