package generics_concurrency

import "fmt"

func MyFilter(input []float64, f func(float64) bool) []float64 {
	var result []float64
	for _, value := range input {
		if f(value) == true {
			result = append(result, value)
		}
	}
	return result
}
func MyGenericFilter[T any](input []T, f func(T) bool) []T {
	var result []T
	for _, value := range input {
		if f(value) == true {
			result = append(result, value)
		}
	}
	return result
}
func FilterAlgorithm() {
	inputFloats := []float64{17.3, 11.1, 9.9, 4.3, 12.6}
	inputIntegers := []int{4, 5, 8, 7, 2, 5}
	result1 := GenericMap[float64, float64](inputFloats, func(n float64) float64 {
		return n * n
	})
	fmt.Printf("Generic map: %v\n", result1)
	greaterThanFive := MyGenericFilter[int](inputIntegers,
		func(i int) bool {
			return i > 5
		},
	)
	fmt.Printf("Generic filter1: %v\n", greaterThanFive)
	res := MyFilter(inputFloats, func(num float64) bool {
		if num <= 10.0 {
			return true
		}
		return false
	})
	fmt.Printf("My Filter result: %v\n", res)
	oddNumbers := MyGenericFilter[int](inputIntegers, func(i int) bool {
		return i%2 == 0
	})
	fmt.Printf("My Generic filter2 result: %v\n", oddNumbers)
	wordsToFilter := []string{"hello", "or", "the", "maybe"}
	lengthGreaterThan3 := MyGenericFilter[string](wordsToFilter, func(s string) bool {
		return len(s) > 3
	})
	fmt.Printf("My Generic filter3 result: %v\n", lengthGreaterThan3)
}
