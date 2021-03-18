package arrays

func findMaxConsecutiveOnes(nums []int) int {
	res := 0
	currentMax := 0
	for _, num := range nums {
		if num == 1 {
			currentMax++
		} else {
			currentMax = 0
		}

		if currentMax > res {
			res = currentMax
		}
	}

	return res
}
