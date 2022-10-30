package generics_concurrency

import (
	"fmt"
	"time"
)

func regularFunction() {
	fmt.Println("Just entered regularFunction()")
	time.Sleep(5 * time.Second)
}
func goroutineFunction() {
	fmt.Println("Just entered goroutineFunction()")
	time.Sleep(10 * time.Second)
	fmt.Println("goroutineFunction finished its work")
}
func MyGoRoutines() {
	go goroutineFunction()
	fmt.Println("In main one line below goroutineFunction()")
	regularFunction()
	fmt.Println("In main one line below regularFunction()")
}
