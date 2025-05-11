// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
)

var i = 0

func incrementing(incrementCh, doneCh chan struct{}) {
	//TODO: increment i 1000000 times
	for j := 0; j < 1000001; j++ {
		incrementCh <- struct{}{}
	}
	doneCh <- struct{}{}
}

func decrementing(decrementCh, doneCh chan struct{}) {
	//TODO: decrement i 1000000 times
	for j := 0; j < 1000000; j++ {
		decrementCh <- struct{}{}
	}
	doneCh <- struct{}{}
}

func counter(incrementCh, decrementCh, doneCh chan struct{}, stopCh chan int) {
	count := 0
	for {
		select {
		case <-incrementCh:
			i++
		case <-decrementCh:
			i--
		case <-doneCh:
			count++
			if count >= 2 {
				stopCh <- i
				return
			}
		}
	}
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(2)

	incrementCh := make(chan struct{})
	decrementCh := make(chan struct{})
	doneCh := make(chan struct{})
	stopCh := make(chan int)

	// TODO: Spawn both functions as goroutines
	go incrementing(incrementCh, doneCh)
	go decrementing(decrementCh, doneCh)
	go counter(incrementCh, decrementCh, doneCh, stopCh)

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	//time.Sleep(500 * time.Millisecond)

	i := <-stopCh
	Println("The magic number is:", i)
}
