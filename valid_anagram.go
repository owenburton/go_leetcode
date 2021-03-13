package main

import "fmt"

// Understand, plan, execute, review
// iterate through the length of s/t
// at each iteration, add the current character of s & t to
// two different maps
// check if the maps are equal
// if yes, return true
// else, return false

func isAnagram(s string, t string) bool {
	sBytes := make(map[byte]int)
	tBytes := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if _, ok := sBytes[s[i]]; ok {
			sBytes[s[i]]++
		} else {
			sBytes[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		if _, ok := tBytes[t[i]]; ok {
			tBytes[t[i]]++
		} else {
			tBytes[t[i]] = 1
		}
	}

	for k, v1 := range sBytes {
		v2, ok := tBytes[k]

		if ok == false || v1 != v2 {
			return false
		}

		delete(tBytes, k)
	}

	if len(tBytes) != 0 {
		return false
	}

	return true
}

func main() {
	s1 := "anagram"
	t1 := "nagaram"

	// isAnagram(s1, t1)
	fmt.Println(isAnagram(s1, t1))

	s2 := "rat"
	t2 := "car"

	// isAnagram(s2, t2)
	fmt.Println(isAnagram(s2, t2))

	s3 := "a"
	t3 := "ab"

	fmt.Println(isAnagram(s3, t3))
}
