package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, val := range nums {
		if i, foundValue := m[target-val]; foundValue {
			return []int{i, idx}
		}
		m[val] = idx
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	var target = 18
	fmt.Println(twoSum(nums, target))
}
