package main

import "fmt"

func removeDuplicates(nums []int) int {
	idx := 0
	for i:=0; i<len(nums); i++ {
		if nums[i] != nums[idx] {
			idx++
			nums[idx] = nums[i]
		}
	}
	return idx+1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}