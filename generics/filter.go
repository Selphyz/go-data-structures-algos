package generics

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
func FilterAlgorithm() {
	input := []float64{17.3, 11.1, 9.9, 4.3, 12.6}
	res := MyFilter(input, func(num float64) bool {
		if num <= 10.0 {
			return true
		}
		return false
	})
	fmt.Println(res)
}
