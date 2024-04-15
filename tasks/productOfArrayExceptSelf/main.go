package main

import "fmt"

/*
[a, b, c, d]
LEFT  [1     a    ab    abc]
RIGHT [bcd  cd     d    1  ]
ANS  [bcd   acd   abd   abc]
*/

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	answer := make([]int, len(nums))
	left[0] = 1
	right[len(nums)-1] = 1

	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
		j := len(nums) - 1 - i
		right[j] = right[j+1] * nums[j+1]
	}

	for i := 0; i < len(nums); i++ {
		answer[i] = left[i] * right[i]
	}
	return answer
}

func main() {
	nums1 := []int{1, 2, 3, 4}
	fmt.Println("Input:", nums1)
	fmt.Println("Output:", productExceptSelf(nums1))

	nums2 := []int{0, 1, 2, 3, 4}
	fmt.Println("\nInput:", nums2)
	fmt.Println("Output:", productExceptSelf(nums2))

	nums3 := []int{-1, 1, 0, -3, 3}
	fmt.Println("\nInput:", nums3)
	fmt.Println("Output:", productExceptSelf(nums3))

	nums4 := []int{1, -1}
	fmt.Println("\nInput:", nums4)
	fmt.Println("Output:", productExceptSelf(nums4))
}
