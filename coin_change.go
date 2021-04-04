package main

import (
	"fmt"
)

/*
Understand, Plan, Execute, Review
Solve this with a breadth first search.
*/

func coinChange(coins []int, amount int) int {
	aQ := []int{amount}
	cache := make(map[int]bool, 0)
	steps := 0
	for {
		size := len(aQ)
		if size == 0 {
			break
		}
		fmt.Println(steps, aQ)

		for {
			if size < 1 {
				break
			}
			size -= 1
			cur := aQ[0]
			aQ = aQ[1:]
			if cur == 0 {
				return steps
			}

			for _, c := range coins {
				remain := cur - c
				if remain < 0 {
					continue
				}

				if _, ok := cache[remain]; ok {
					continue
				} else {
					aQ = append(aQ, remain)
					cache[remain] = true
				}
			}
		}

		steps += 1
	}

	return -1
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11

	// coins := []int{2}
	// amount := 3

	fmt.Println(coinChange(coins, amount))
}

// Solve this shit using recursion
