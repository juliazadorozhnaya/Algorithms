package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	// Шаг 1: Подсчет частоты каждого элемента
	allElem := make(map[int]int)
	for _, value := range nums {
		allElem[value]++
	}

	// Шаг 2: создаем массив для списка частот
	bucket := make([][]int, len(nums)+1)
	for num, freq := range allElem {
		bucket[freq] = append(bucket[freq], num)
	}

	// Шаг 3: Идем по бакету с конца
	result := make([]int, 0, k)
	for freq := len(bucket) - 1; len(result) < k && freq >= 0; freq-- {
		result = append(result, bucket[freq]...) //для плоского среза многоточие
	}
	return result
}

func main() {
	nums1 := []int{1, 1, 1, 2, 2, 3}
	k1 := 2
	fmt.Println(topKFrequent(nums1, k1)) // Output: [1, 2]

	nums2 := []int{1}
	k2 := 1
	fmt.Println(topKFrequent(nums2, k2)) // Output: [1]
}
