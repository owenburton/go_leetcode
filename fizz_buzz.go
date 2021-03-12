package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	// Understand, plan, execute, review
	// instantiate an empty list of length input number
	// iterate through 1 to the input number
	// if the current iteration number is divisible by 15,
	// change the current index element in answer array to "FizzBuzz"
	// if the current num divisible by 5,
	// change the current index element in answer array to "Buzz"
	// if the current num divisible by 3,
	// change the current index element in answer array to "Fizz"
	// else
	// convert the current int to a string and put it in the answer array

	answer := make([]string, n)
	for num := 1; num <= n; num++ {
		switch {
		case num%15 == 0:
			answer[num-1] = "FizzBuzz"
		case num%5 == 0:
			answer[num-1] = "Buzz"
		case num%3 == 0:
			answer[num-1] = "Fizz"
		default:
			answer[num-1] = strconv.Itoa(num)
		}
	}

	return answer
}

func main() {
	input := 15

	fmt.Println(fizzBuzz(input))
}
