package main

import (
	"fmt"
)

func reverse(x int) int {
	const (
		int32Max = 2147483647  // 2^31 - 1
		int32Min = -2147483648 // -2^31
	)
	result := 0

	for x != 0 {
		pop := x % 10
		x /= 10

		if result > int32Max/10 || (result == int32Max/10 && pop > 7) {
			return 0
		}
		if result < int32Min/10 || (result == int32Min/10 && pop < -8) {
			return 0
		}

		result = result*10 + pop
	}
	return result
}

func main() {
	fmt.Println(reverse(123))  // Output: 321
	fmt.Println(reverse(-123)) // Output: -321
	fmt.Println(reverse(120))  // Output: 21
}
