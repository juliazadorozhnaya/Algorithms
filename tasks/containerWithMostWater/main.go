package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	var maxArea int

	for left < right {
		h := min(height[left], height[right])
		l := right - left
		currentArea := h * l

		if currentArea > maxArea {
			maxArea = currentArea
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
