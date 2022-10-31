package sort_search

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

const size = 1_000_000_000

var data []float64

func isSegmentSorted(data []float64, a, b int, ch chan<- bool) {
	// Generates boolean value put into ch
	for i := a + 1; i < b; i++ {
		if data[i] < data[i-1] {
			ch <- false
		}
	}
	ch <- true
}
func isSorted2(data []float64) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}
func IsSortedConcurrent(data []float64) bool {
	ch := make(chan bool)
	numSegments := runtime.NumCPU()
	segmentSize := int(float64(len(data)) / float64(numSegments))
	// Launch numSegments goroutines
	for index := 0; index < numSegments; index++ {
		go isSegmentSorted(data, index*segmentSize,
			index*segmentSize+segmentSize, ch)
	}
	num := 0 // Completed goroutines
	for {
		select {
		case value := <-ch:
			if value == false {
				return false
			}
			num += 1
			if num == numSegments { // All goroutiines have completed
				return true
			}
		}
	}
	return true
}
func CheckSorted() {
	data = make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}
	data2 := make([]float64, size)
	// Create a sorted sequence of numbers
	for i := 0; i < size; i++ {
		data2[i] = float64(2 * i)
	}
	start := time.Now()
	result := isSorted2(data)
	elapsed := time.Since(start)
	fmt.Println("\nSorted: ", result)
	fmt.Println("elapsed using sorted2", elapsed)
	start = time.Now()
	result = isSorted2(data2)
	elapsed = time.Since(start)
	fmt.Println("Sorted: ", result)
	fmt.Println("elapsed using sorted2:", elapsed)
	start = time.Now()
	result = IsSortedConcurrent(data)
	elapsed = time.Since(start)
	fmt.Println("\nSorted: ", result)
	fmt.Println("elapsed using concurrent sorted3", elapsed)
	start = time.Now()
	result = IsSortedConcurrent(data2)
	elapsed = time.Since(start)
	fmt.Println("Sorted: ", result)
	fmt.Println("elapsed using concurrent sorted3:", elapsed)
}
