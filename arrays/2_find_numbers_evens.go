package arrays

func findNumbers(nums []int) int {
	count := 0
	for _, num := range nums {
		if recursionNum(num) % 2 == 0 {
			count++
		}
	}
	return count
}

func recursionNum (num int) int {
	if num < 10 {
		return 1
	} else {
		return 1 + recursionNum(num/10)
	}
}
