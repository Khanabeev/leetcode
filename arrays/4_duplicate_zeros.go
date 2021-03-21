package arrays

func duplicateZeros1(arr []int) []int {
	count := len(arr)

	for i, num := range arr {
		if num == 0 {
			arr = insert(arr, 0, i+1)
		}
	}

	return arr[:count]
}

func insert(array []int, element int, i int) []int {
	return append(array[:i], append([]int{element}, array[i:]...)...)
}

func duplicateZeros(arr []int) []int {
	count := len(arr) - 1
	possibleDups := 0

	for i := 0; i <= count-possibleDups; i++ {

		if arr[i] == 0 {
			if i == count-possibleDups {
				arr[count] = 0
				count -= 1
				break
			}
			possibleDups++
		}
	}

	last := count - possibleDups

	for i := last; i >= 0; i-- {
		if arr[i] == 0 {
			arr[i+possibleDups] = 0
			possibleDups--
			arr[i+possibleDups] = 0
		} else {
			arr[i+possibleDups] = arr[i]
		}
	}

	return arr
}
