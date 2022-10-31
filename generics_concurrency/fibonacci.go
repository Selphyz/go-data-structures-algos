package generics_concurrency

import "fmt"

func fibonacci(c chan<- int, quit <-chan bool) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y // Generates the sequence
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func Fibonacci() {
	c := make(chan int)
	quit := make(chan bool)
	go func() {
		for i := 0; i < 30; i++ {
			fmt.Println(<-c)
		}
		quit <- true
	}()
	fibonacci(c, quit)
}
