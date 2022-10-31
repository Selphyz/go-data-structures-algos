package generics_concurrency

import (
	"fmt"
	"time"
)

var quit chan bool

func pingGen(c chan<- string) {
	for i := 0; i < 5; i++ {
		c <- "ping"
	}
}
func pongGen(c chan<- string) {
	for i := 0; i < 5; i++ {
		c <- "pong"
	}
}
func outputLoop(c <-chan string) {
	for {
		time.Sleep(time.Second * 1)
		select {
		case value := <-c:
			fmt.Println(value)
		case <-time.After(3 * time.Second):
			fmt.Println("Program timed out.")
			quit <- true
		}
	}
}
func ChannelEnd() {
	quit = make(chan bool)
	c := make(chan string)
	go pingGen(c)
	go pongGen(c)
	go outputLoop(c)
	<-quit
}
