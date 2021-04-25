package main

import "fmt"

// Understand, plan, execute, review.
// Iterate backwards through the array.
// Have a variable called previous.
// Make previous the value at the last index of arr.
// Iterate backwards through the array.
// At each index compare previous and the current value.
// Make the current value into previous and
// Change previous to the current value if it was greater.
// After done iterating, change the last value in arr to -1.

func replaceElements(arr []int) []int {
	prev := -1
	for i:=len(arr)-1; i>=0; i-- {
		temp := arr[i]
		arr[i] = prev
		if temp > prev {
			prev = temp
		}
	}
	return arr
}

func main() {
	arr := []int{17,18,5,4,6,1}
	fmt.Println(arr)
	fmt.Println(replaceElements(arr))
}