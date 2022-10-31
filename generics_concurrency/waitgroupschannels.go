package generics_concurrency

import (
	"fmt"
	"sync"
	"time"
)

var wgp sync.WaitGroup

func pingGenerator(c chan string) {
	defer wgp.Done()
	for i := 0; i < 5; i++ {
		c <- "ping"
		time.Sleep(time.Second * 1)
	}
}
func output(c chan string) {
	for {
		select {
		case value := <-c:
			fmt.Println(value)
		case <-time.After(3 * time.Second):
			fmt.Println("Program timed out.")
			wg.Done()
		}
	}
}
func PingCall() {
	c := make(chan string)
	wgp.Add(2)
	go pingGenerator(c)
	go output(c)
	wgp.Wait()
}
