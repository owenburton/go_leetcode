package main

import "fmt"

// Understand, plan, execute, review
// Copy over first m values from nums1 to temp slice.
// Iterate through the index values of nums1.
// If m & n do not equal zero, then
// if m greater than or equal to n, then insert m-1 value of temp into
// the current index of nums1.
// If n is greater, then insert n-1 value of nums2 into nums1 current index.
// If either m or n is equal to zero,
// iterate m times and insert the rest of temp into nums1, then
// iterate n times and insert the rest of nums2 into nums1.

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := len(nums1) - 1
	for {
		if m>0 && n>0 {
			if nums1[m-1] >= nums2[n-1] {
				nums1[i] = nums1[m-1]
				m--
			} else {
				nums1[i] = nums2[n-1]
				n--
			}
			i--
		} else {
			break
		}
	}

	for {
		if n<1 {
			break
		}
		nums1[i] = nums2[n-1]
		i--
		n--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	// nums1 := []int{0}
	// m := 0
	// nums2 := []int{1}
	// n := 1

	merge(nums1, m, nums2, n)

	fmt.Println(nums1)
}