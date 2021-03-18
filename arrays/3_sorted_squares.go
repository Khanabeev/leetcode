package arrays

import "sort"

func sortedSquares(nums []int) []int {
	var arr []int
	for _, num := range nums{
		res := num * num
		arr = append(arr,res)
	}
	sort.Ints(arr)
	return arr
}
