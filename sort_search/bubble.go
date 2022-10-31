package sort_search

import (
	"fmt"
)

type bOrdered interface {
	~float64 | ~int | ~string
}

func bubblesort[T bOrdered](data []T) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func BubbleSort() {
	numbers := []float64{3.5, -2.4, 12.8, 9.1}
	names := []string{"Zachary", "John", "Moe", "Jim", "Robert"}
	bubblesort[float64](numbers)
	fmt.Printf("%v\n", numbers)
	bubblesort[string](names)
	fmt.Printf("%v\n", names)
}
