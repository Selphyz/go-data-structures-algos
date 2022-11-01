package sort_search

import (
	"fmt"
	"math"
	"time"
)

const sizeh = 100_000

func BubbleComparison() {
	data := make([]float64, sizeh)
	for i := 0; i < sizeh; i++ {
		data[i] = math.Sin(float64(i * i))
	}
	start := time.Now()
	quicksort[float64](data, 0, len(data)-1)
	elapsed := time.Since(start)
	fmt.Println("Elapsed sort time for sine wave using quicksort: ",
		elapsed)
	data = make([]float64, sizeh)
	for i := 0; i < sizeh; i++ {
		data[i] = math.Sin(float64(i * i))
	}
	start = time.Now()
	bubblesort[float64](data)
	elapsed = time.Since(start)
	fmt.Println("Elapsed sort time for sine wave using bubblesort: ", elapsed)
}
