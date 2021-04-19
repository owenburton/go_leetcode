package main

import "fmt"

func removeElement(nums []int, val int) int {
	idx := 0
	for _, v := range nums {
		if v!=val {
			nums[idx] = v
			idx++
		}
	}
	return idx
}

func main() {
	nums := []int{0,1,2,2,3,0,4,2}
	val := 2
	
	fmt.Println(removeElement(nums, val))
	fmt.Println(nums)
}