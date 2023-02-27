package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	lower := ""

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			lower += string([]byte{s[i]})
		}
		if s[i] >= '0' && s[i] <= '9' {
			lower += string([]byte{s[i]})
		}
	}

	for i := 0; i < len(lower)/2; i++ {
		if lower[i] != lower[len(lower)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	str := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(str))
}
