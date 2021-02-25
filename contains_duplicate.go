package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	cache := make(map[int]int)
	for _, v := range nums {
		if _, ok := cache[v]; ok {
			return true
		}
		cache[v] = 0
	}
	return false
}

func main() {
	input := []int{1, 2, 3, 1}

	fmt.Println(containsDuplicate(input))
}
