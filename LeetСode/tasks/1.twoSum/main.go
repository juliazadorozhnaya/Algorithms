package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {
		if v, found := m[target-num]; found {
			return []int{v, idx}
		}
		m[num] = idx // добавим в мапу элемент из nums под таким же индексом
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	var target = 18
	fmt.Println(twoSum(nums, target))
}
