package main

import "fmt"

// Understand, Plan, Execute, Review
// Instantiate a slice, "answer", with [1] in it.
// Loop for one less than numRows times.
// For each iteration, get the current length of the answer slice.
// For each iteration, create a slice containing 1 that will be appended to answer.
// For each iteration, loop the current length of answer slice times.
// For each nested iteration, check if a 1 needs to be appended on the right side.
// If not, append the sum of the above two numbers.
// For each outer iteration, end by appending the constructed sublice.
// Once the outer loop is finished, return answer slice.

func generate(numRows int) [][]int {
	answer := [][]int{{1}}
	for i := 0; i < numRows-1; i++ {
		answerLen := len(answer)
		subSlice := []int{1}
		for j := 0; j < answerLen; j++ {
			if j == answerLen-1 {
				subSlice = append(subSlice, 1)
				break
			}
			aboveSum := answer[answerLen-1][j] + answer[answerLen-1][j+1]
			subSlice = append(subSlice, aboveSum)
		}
		answer = append(answer, subSlice)
	}
	return answer
}

func main() {
	numRows := 7
	fmt.Println(generate(numRows))
}
