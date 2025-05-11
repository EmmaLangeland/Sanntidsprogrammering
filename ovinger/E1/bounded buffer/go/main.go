package main

import (
	"fmt"
	"time"
)

func producer(bufferedCh chan int) {

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[producer]: pushing %d\n", i)

		// TODO: push real value to buffer
		bufferedCh <- i
	}

}

func consumer(bufferedCh chan int) {

	time.Sleep(1 * time.Second)
	for {
		//TODO: get real value from buffer
		i := <-bufferedCh

		fmt.Printf("[consumer]: %d\n", i)
		time.Sleep(50 * time.Millisecond)
	}

}

func main() {

	// TODO: make a bounded buffer
	bufferedCh := make(chan int, 5)

	go consumer(bufferedCh)
	go producer(bufferedCh)

	select {}
}
