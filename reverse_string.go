package main

import "fmt"

func reverseString(s []byte) {
	// Understand, plan, execute, review.
	// Iterate through the bytes in the slice.
	// At each iteration, switch the current index with it's mirror across
	// half the slice.
	for i, v := range s {
		if i >= len(s)/2 {
			return
		}
		temp := s[len(s)-1-i]
		s[len(s)-1-i] = v
		s[i] = temp
	}

}

func main() {
	// input := []byte{'h', 'e', 'l', 'l', 'o'}
	inputSlice := []string{
		"A", " ", "m", "a", "n", ",", " ", "a", " ", "p", "l", "a", "n",
		",", " ", "a", " ", "c", "a", "n", "a", "l", ":", " ", "P", "a",
		"n", "a", "m", "a"}

	inputStr := ""
	for _, v := range inputSlice {
		inputStr += v
	}
	inputBytes := []byte(inputStr)

	reverseString(inputBytes)

	fmt.Println("My answer:")
	fmt.Println(string(inputBytes))

	correctSlice := []string{
		"a", "m", "a", "n", "a", "P", " ", ":", "l", "a", "n",
		"a", "c", " ", "a", " ", ",", "n", "a", "l", "p", " ",
		"a", " ", ",", "n", "a", "m", " ", "A"}
	correctStr := ""
	for _, v := range correctSlice {
		correctStr += v
	}
	correctBytes := []byte(correctStr)

	fmt.Println("Correct answer:")
	fmt.Println(string(correctBytes))

	if string(inputBytes) == string(correctBytes) {
		fmt.Println("You have the correct answer!")
	}
}
