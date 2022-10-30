package generics

import "fmt"

func MyMap(input []int, f func(int) int) []int {
	result := make([]int, len(input))
	for index, val := range input {
		result[index] = f(val)
	}
	return result
}
func GenericMap[T1, T2 any](input []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(input))
	for index, value := range input {
		result[index] = f(value)
	}
	return result
}
func MapsAlgorithm() {
	slice := []int{1, 3, 4, 7, 6}
	result := MyMap(slice, func(i int) int {
		return i * i
	})
	fmt.Println(result)
	result2 := GenericMap(slice, func(i int) int {
		return i * 5
	})
	fmt.Println(result2)
}
