package sort_search

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const csize = 50_000_000
const threshold = 5000

type cOrdered interface {
	~float64 | ~int | ~string
}

func InsertSort[T cOrdered](data []T) {
	i := 1
	for i < len(data) {
		h := data[i]
		j := i - 1
		for j >= 0 && h < data[j] {
			data[j+1] = data[j]
			j -= 1
		}
		data[j+1] = h
		i += 1
	}
}
func Partition[T cOrdered](data []T) int {
	data[len(data)/2], data[0] = data[0], data[len(data)/2]
	pivot := data[0]
	mid := 0
	i := 1
	for i < len(data) {
		if data[i] < pivot {
			mid += 1
			data[i], data[mid] = data[mid], data[i]
		}
		i += 1
	}
	data[0], data[mid] = data[mid], data[0]
	return mid
}
func IsSorted[T cOrdered](data []T) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}
func ConcurrentQuicksort[T cOrdered](data []T, wg *sync.WaitGroup) {
	for len(data) >= 30 {
		mid := Partition(data)
		var portion []T
		if mid < len(data)/2 {
			portion = data[:mid]
			data = data[mid+1:]
		} else {
			portion = data[mid+1:]
			data = data[:mid]
		}
		if len(portion) > threshold {
			wg.Add(1)
			go func(data []T) {
				defer wg.Done()
				ConcurrentQuicksort(data, wg)
			}(portion)
		} else {
			ConcurrentQuicksort(portion, wg)
		}
	}
	InsertSort(data)
}
func QSort[T cOrdered](data []T) {
	var wg sync.WaitGroup
	ConcurrentQuicksort(data, &wg)
	wg.Wait()
}
func cpartition[T cOrdered](data []T, low, high int) int {
	var pivot = data[low]
	var i = low
	var j = high
	for i < j {
		for data[i] <= pivot && i < high {
			i++
		}
		for data[j] > pivot && j > low {
			j--
		}
		if i < j {
			data[i], data[j] = data[j], data[i]
		}
	}
	data[low] = data[j]
	data[j] = pivot
	return j
}
func cquicksort[T cOrdered](data []T, low, high int) {
	if low < high {
		var pivot = cpartition(data, low, high)
		cquicksort(data, low, pivot)
		cquicksort(data, pivot+1, high)
	}
}
func ConcurrentQSort() {
	data := make([]float64, csize)
	for i := 0; i < csize; i++ {
		data[i] = 100.0 * rand.Float64()
	}
	data2 := make([]float64, csize)
	copy(data2, data)
	start := time.Now()
	QSort[float64](data)
	elapsed := time.Since(start)
	fmt.Println("Elapsed time for concurrent quicksort = ", elapsed)
	fmt.Println("Is sorted: ", IsSorted(data))
	start = time.Now()
	cquicksort(data2, 0, len(data2)-1)
	elapsed = time.Since(start)
	fmt.Println("Elapsed time for regular quicksort = ", elapsed)
	fmt.Println("Is sorted: ", IsSorted(data2))
}
