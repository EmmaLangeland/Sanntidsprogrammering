// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i = 1

// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
func incrementing() {
	//TODO: increment i 1000000 times
	for j := 0; j < 1000000; j++ {
		i++
	}
}

func decrementing() {
	//TODO: decrement i 1000000 times
	for j := 0; j < 1000001; j++ {
		i--
	}
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	// GOMAXPROCS set the number of CPU threads. Using runtime.GOMAXPROCS to ensure a 1:1 mapping of goroutines to CPU cores (or threads).
	// If set to 1, only one CPU will be used and the result will always end up as 0 due to no overlap. But it will work slow.
	runtime.GOMAXPROCS(2)

	// TODO: Spawn both functions as goroutines
	go incrementing()
	go decrementing()

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	time.Sleep(500 * time.Millisecond)
	Println("GO The magic number is:", i)
}
