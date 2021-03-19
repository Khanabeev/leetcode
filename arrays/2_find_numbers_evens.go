package arrays

func findNumbers2(nums []int) int {
	count := 0
	for _, num := range nums {
		if recursionNum2(num) % 2 == 0 {
			count++
		}
	}
	return count
}

func recursionNum2 (num int) int {
	if num < 10 {
		return 1
	} else {
		return 1 + recursionNum(num/10)
	}
}
