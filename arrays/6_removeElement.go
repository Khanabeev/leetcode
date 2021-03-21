package arrays

func removeElement(nums []int, val int) int {
	index := 0
	for _, i := range nums {
		if i != val {
			nums[index] = i
			index++
		}
	}

	return len(nums[:index])
}

// do not modify original array
func findAndDelete(s []int, itemToDelete int) []int {
	var new []int
	index := 0
	for _, i := range s {
		if i != itemToDelete {
			new[index] = i
			index++
		}
	}
	return new[:index]
}
