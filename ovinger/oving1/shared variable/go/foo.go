// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i = 1

func do(increase_chan, decrease_chan, done_chan, return_chan chan int) {
	donecounter := 0
	for {
		select {
		case <-increase_chan:
			i++
		case <-decrease_chan:
			i--
		case <-done_chan:
			donecounter++
			if donecounter >= 2 {
				return_chan <- i
				return
			}
		}
	}
}

// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
func incrementing(increase_chan, done_chan chan int) {
	//TODO: increment i 1000000 times
	for j := 0; j < 1000000; j++ {
		increase_chan <- 1
	}
	done_chan <- 1
}

func decrementing(decrease_chan, done_chan chan int) {
	//TODO: decrement i 1000000 times
	for j := 0; j < 1000001; j++ {
		decrease_chan <- 1
	}
	done_chan <- 1
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	// GOMAXPROCS set the number of CPU threads. Using runtime.GOMAXPROCS to ensure a 1:1 mapping of goroutines to CPU cores (or threads).
	// If set to 1, only one CPU will be used and the result will always end up as 0 due to no overlap. But it will work slow.
	runtime.GOMAXPROCS(2)

	increase_chan := make(chan int)
	decrease_chan := make(chan int)
	done_chan := make(chan int)
	return_chan := make(chan int)

	go incrementing(increase_chan, done_chan)
	go decrementing(decrease_chan, done_chan)
	go do(increase_chan, decrease_chan, done_chan, return_chan)

	// TODO: Spawn both functions as goroutines
	//go incrementing()
	//go decrementing()

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	time.Sleep(500 * time.Millisecond)
	Println("GO The magic number is:", i)

}
