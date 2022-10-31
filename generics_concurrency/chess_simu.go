package generics_concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

var quitchess chan bool

func ChessSimulation() {
	rand.Seed(time.Now().UnixNano())
	move := make(chan int)
	quitchess = make(chan bool)
	go player("Bobby Fischer", move)
	go player("Boris Spassky", move)

	move <- 1
	<-quitchess
}

func player(name string, move chan int) {
	for {
		turn := <-move
		n := rand.Intn(100)
		if n <= 5 && turn >= 5 {
			fmt.Printf("Player %s was check mated and loses!", name)
			quitchess <- true
			return
		}
		fmt.Printf("Player %s has moved. Turn %d.\n", name, turn)
		turn++
		time.Sleep(1 * time.Second)
		move <- turn
	}
}
