package controller

import (
	"sort"
	"sync"
	"time"
)

func ProcessSequential(subArrays [][]int) ([][]int, int64) {
	startTime := time.Now()

	var sortedArrays [][]int
	for _, subArray := range subArrays {
		sortedSubArray := make([]int, len(subArray))
		copy(sortedSubArray, subArray)
		sort.Ints(sortedSubArray)
		sortedArrays = append(sortedArrays, sortedSubArray)
	}

	elapsedTime := time.Since(startTime).Nanoseconds()
	return sortedArrays, elapsedTime
}

func sortArray(subArray []int, ch chan<- []int, wg *sync.WaitGroup) {
	defer wg.Done()

	sortedSubArray := make([]int, len(subArray))
	copy(sortedSubArray, subArray)
	sort.Ints(sortedSubArray)

	ch <- sortedSubArray
}

func ProcessConcurrent(subArrays [][]int) ([][]int, int64) {
	startTime := time.Now()

	var wg sync.WaitGroup
	ch := make(chan []int)

	sortedArrays := make([][]int, len(subArrays))

	for _, subArray := range subArrays {
		wg.Add(1)
		go sortArray(subArray, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	i := 0

	for result := range ch {
		sortedArrays[i] = result
		i = i + 1
	}

	elapsedTime := time.Since(startTime).Nanoseconds()
	return sortedArrays, elapsedTime
}
