package sort_search

import (
	"fmt"
	"time"
)

const binsize = 100_000_000

type binOrdered interface {
	~float64 | ~int | ~string
}

func binarySearch[T binOrdered](slice []T, target T) bool {
	low := 0
	high := len(slice) - 1
	for low <= high {
		median := (low + high) / 2
		if slice[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low == len(slice) || slice[low] != target {
		return false
	}
	return true
}
func BinarySearch() {
	data := make([]float64, binsize)
	for i := 0; i < binsize; i++ {
		data[i] = float64(i)
	}
	start := time.Now()
	result := binarySearch[float64](data, -10.0)
	elapsed := time.Since(start)
	fmt.Println("Time to search slice of 100_000_000 floats using binarySearch = ", elapsed)
	fmt.Println("Result of search is ", result)
	start = time.Now()
	result = binarySearch[float64](data, float64(binsize/2))
	elapsed = time.Since(start)
	fmt.Println("Time to search slice of 100_000_000 floats using binarySearch = ", elapsed)
	fmt.Println("Result of search is ", result)
}
