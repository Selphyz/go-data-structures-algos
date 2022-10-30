package main

import "fmt"

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

type Student struct {
	Name string
	ID   int
	age  float64
}

func (s Student) String() string {
	return fmt.Sprintf("%s %0.2f %d", s.Name, s.age, s.ID)
}
func addStudent[T Stringer](students []T, student T) []T {
	return append(students, student)
}
func main() {
	var students []String
	result := addStudent[String](students, "Michael")
	result = addStudent[String](result, "Jenny")
	result = addStudent[String](result, "Elaine")
	fmt.Println(result)
	var student1 []Integer
	result1 := addStudent[Integer](student1, 155)
	result1 = addStudent[Integer](result1, 112)
	result1 = addStudent[Integer](result1, 120)
	fmt.Println(result1)
	var students2 []Student // Empty slice
	result2 := addStudent(students2, Student{"John", 213, 17.5})
	result2 = addStudent(result2, Student{"James", 111, 18.75})
	result2 = addStudent(result2, Student{"Marsha", 110, 16.25})
	fmt.Println(result2)
}
