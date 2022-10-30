package generics_concurrency

import (
	"fmt"
	"sort"
)

type Stringer = interface {
	String() string
}
type Integer int
type String string

func (i Integer) String() string {
	return fmt.Sprintf("%d", i)
}
func (s String) String() string {
	return string(s)
}

type Ordered interface {
	~int | ~float64 | ~string
}
type Student struct {
	Name string
	ID   int
	Age  float64
}

func (s Student) String() string {
	return fmt.Sprintf("%s %0.2f %d", s.Name, s.Age, s.ID)
}
func addStudent[T Stringer](students []T, student T) []T {
	return append(students, student)
}

// OrderedSlice prepared to be ordered
type OrderedSlice[T Ordered] []T

func (s OrderedSlice[T]) Len() int {
	return len(s)
}
func (s OrderedSlice[T]) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s OrderedSlice[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// SortType prepared to be ordered
type SortType[T any] struct {
	slice   []T
	compare func(T, T) bool
}

func (s SortType[T]) Len() int {
	return len(s.slice)
}
func (s SortType[T]) Less(i, j int) bool {
	return s.compare(s.slice[i], s.slice[j])
}
func (s SortType[T]) Swap(i, j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}
func PerformSort[T any](slice []T, compare func(T, T) bool) {
	sort.Sort(SortType[T]{slice, compare})
}
func SortAlgorthm() {
	var students []String
	result := addStudent[String](students, "Michael")
	result = addStudent[String](result, "Jenny")
	result = addStudent[String](result, "Elaine")
	sort.Sort(OrderedSlice[String](result))
	fmt.Printf("%v\n", result)
	var student1 []Integer
	result1 := addStudent[Integer](student1, 155)
	result1 = addStudent[Integer](result1, 112)
	result1 = addStudent[Integer](result1, 120)
	sort.Sort(OrderedSlice[Integer](result1))
	fmt.Printf("%v\n", result1)
	var students2 []Student // Empty slice
	result2 := addStudent(students2, Student{"John", 213, 17.5})
	result2 = addStudent(result2, Student{"James", 111, 18.75})
	result2 = addStudent(result2, Student{"Marsha", 110, 16.25})
	PerformSort[Student](result2, func(s1, s2 Student) bool {
		return s1.Age < s2.Age
	})
	fmt.Printf("%v\n", result2)
}
