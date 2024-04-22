package main

import "fmt"

func romanToInt(s string) int {
	romanToIntMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romanToIntMap[rune(s[i])] < romanToIntMap[rune(s[i+1])] {
			result -= romanToIntMap[rune(s[i])]
		} else {
			result += romanToIntMap[rune(s[i])]
		}
	}

	return result
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("MCMXCIV")) //1994
}
