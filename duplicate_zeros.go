package main

import "fmt"

// Understand, plan, execute, review.
//
func duplicateZeros(arr []int) {
	temp := make([]int, len(arr))
	j := 0
	for i:=0; j<len(temp); i++ {
		temp[j] = arr[i]
		if arr[i] == 0 {
			j++
			if j==len(arr) {
				break
			}
			temp[j] = 0
		}
		j++
	}

	copy(arr, temp)
}

func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr)
	fmt.Println(arr)
}
