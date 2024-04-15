package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	str1 := make(map[rune]int)
	str2 := make(map[rune]int)

	for _, char := range s {
		str1[char]++
	}

	for _, char := range t {
		str2[char]++
	}

	for key2, value2 := range str2 {
		if str1[key2] != value2 {
			return false
		}
	}

	return true
}

func main() {
	// Первый пример
	s1 := "a"
	t1 := "ab"
	result1 := isAnagram(s1, t1)
	fmt.Printf("Is '%s' an anagram of '%s'? %v\n", s1, t1, result1)
}
