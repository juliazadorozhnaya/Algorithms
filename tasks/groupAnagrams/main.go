package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	mapStr := make(map[string][]string)

	for _, str := range strs {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		mapStr[string(runes)] = append(mapStr[string(runes)], str)
	}

	result := make([][]string, 0, len(mapStr))
	for _, group := range mapStr {
		result = append(result, group)
	}
	return result
}

func main() {
	st := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Print(groupAnagrams(st))
}
