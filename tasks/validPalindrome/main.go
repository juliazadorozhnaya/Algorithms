package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		for left < right && !unicode.IsLetter(rune(s[left])) && !unicode.IsNumber(rune(s[left])) {
			left++
		}
		for left < right && !unicode.IsLetter(rune(s[right])) && !unicode.IsNumber(rune(s[right])) {
			right--
		}
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	examples := []string{"Race car", "12321", "Hello, World!", "A man, a plan, a canal, Panama"}

	for _, example := range examples {
		if isPalindrome(example) {
			fmt.Printf("%s is a palindrome.\n", example)
		} else {
			fmt.Printf("%s is not a palindrome.\n", example)
		}
	}
}
