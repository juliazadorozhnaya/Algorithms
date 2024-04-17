package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	start := nums[0]
	var rangeSlice []string
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 1 {
			if start == nums[i-1] {
				rangeSlice = append(rangeSlice, strconv.Itoa(start))
			} else {
				rangeSlice = append(rangeSlice, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i-1]))
			}
			start = nums[i]
		}
	}

	if start == nums[len(nums)-1] {
		rangeSlice = append(rangeSlice, strconv.Itoa(start))
	} else {
		rangeSlice = append(rangeSlice, strconv.Itoa(start)+"->"+strconv.Itoa(nums[len(nums)-1]))
	}

	return rangeSlice
}

func main() {
	nums := []int{0, 2, 3, 4, 6, 8, 9, 10}
	fmt.Println(summaryRanges(nums))
}

/*
Input: nums = [0,1,2,4,5,7]
Output: ["0->2","4->5","7"]
Explanation: The ranges are:
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"
*/
