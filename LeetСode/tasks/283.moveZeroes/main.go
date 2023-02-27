package main

import "fmt"

func main() {
	nums := []int{8, 0, 0, 1}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	p := 0
	for i, _ := range nums {
		if nums[i] != 0 {
			fmt.Println(i, p)
			nums[p], nums[i] = nums[i], nums[p]
			p++
		}
	}

}
