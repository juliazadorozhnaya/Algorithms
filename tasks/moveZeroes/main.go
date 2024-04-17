package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)

	nums = []int{0}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	lastNoZeroEl := 0

	for _, value := range nums {
		if value != 0 {
			nums[lastNoZeroEl] = value
			lastNoZeroEl++
		}
	}

	for i := lastNoZeroEl; i < len(nums); i++ {
		nums[i] = 0
	}
}
