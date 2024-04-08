package main

import "fmt"

func moveZeroes(nums []int) {
	p := 0
	for i, _ := range nums {
		if nums[i] != 0 {
			nums[p], nums[i] = nums[i], nums[p]
			p++
		}
	}
}

func main() {
	nums := []int{0}
	moveZeroes(nums)
	fmt.Println(nums)
}
