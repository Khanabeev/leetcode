package arrays

import "fmt"

func removeDuplicates(nums []int) int {
	n, d := RemoveDuplicates(nums)
	fmt.Println(d)
	return len(n)
}

func RemoveDuplicates(s []int) ([]int, []int) {
	encountered := make(map[int]struct{})
	result := make([]int, 0)
	duplicate := make([]int, 0)
	for _, v := range s {
		if _, ok := encountered[v]; ok {
			duplicate = append(duplicate, v)
			continue
		} else {
			encountered[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result, duplicate
}

func RemoveDuplicates2(elements []int) []int {
	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}
