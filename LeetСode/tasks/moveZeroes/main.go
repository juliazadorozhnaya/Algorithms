package main

import "fmt"

func main() {
	nums := []int{8, 0, 1}
	moveZeroes(nums)
	fmt.Print(nums)
}

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	p := 0
	for i, _ := range nums {
		if nums[i] != 0 {
			nums[p], nums[i] = nums[i], nums[p]
			p++
		}
	}

}
