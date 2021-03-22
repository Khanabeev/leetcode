package arrays

import "sort"

func removeDuplicates(nums []int) int {
	sort.Ints(nums)
	return len(unique(nums))
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
