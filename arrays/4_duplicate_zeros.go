package arrays

import (
	"fmt"
)

func duplicateZeros(arr []int) []int {
	count := len(arr)
	res := make([]int, count)
	copy(res, arr)

	//1,0,0,3,0,4,5,0
	for i, num := range res {
		if num == 0 {
			arr = append(arr, 0)
			copy(arr[i+1:], arr[i:])
		}
	}
	fmt.Println("arr=", arr, "init=", res, "got=", arr[:count])

	return arr[:count]
}
