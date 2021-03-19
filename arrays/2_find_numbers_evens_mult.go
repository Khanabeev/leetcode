package arrays

import (
	"sync"
)

type CountEven struct {
	mu    sync.Mutex
	count int
}

func NewCountEven() *CountEven {
	return &CountEven{}
}

func (r *CountEven) find(nums []int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, num := range nums {
		if recursionNum(num)%2 == 0 {
			r.count++
		}
	}
}

func findNumbers(nums []int) int {
	countEven := NewCountEven()
	numCPU := 4
	chunkSize := (len(nums) + numCPU - 1) / numCPU

	var wg sync.WaitGroup
	for i := 0; i < len(nums); i += chunkSize {
		wg.Add(1)

		end := i + chunkSize
		if end > len(nums) {
			end = len(nums)
		}

		go func(w *sync.WaitGroup, i, end int) {
			countEven.find(nums[i:end])
			w.Done()
		}(&wg, i, end)

		wg.Wait()
	}

	return countEven.count
}

func recursionNum(num int) int {
	if num < 10 {
		return 1
	} else {
		return 1 + recursionNum(num/10)
	}
}
