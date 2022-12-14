package generics_concurrency

import (
	"fmt"
	"sync"
)

const number = 1000

var countValue int
var m sync.Mutex

func Mutex() {
	var wg sync.WaitGroup
	wg.Add(number)
	for i := 0; i < number; i++ {
		go func() {
			m.Lock()
			countValue++
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("\ncountValue = %d", countValue)
}
