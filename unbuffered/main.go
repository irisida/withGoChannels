package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printCounts(label string, count chan int) {
	p := fmt.Println
	pf := fmt.Printf
	// defer the done
	defer wg.Done()

	for {
		// receiving messages from channel
		val, ok := <-count

		if !ok {
			p("Not OK -> The channel was closed")
			return
		}

		pf("Count: %d received from %s \n", val, label)

		if val == 10 {
			pf("The channel is being closed by %s \n", label)
			close(count) // close the channel
			return
		}
		val++
		// send count back to the other goroutine
		count <- val
	}
}

func main() {
	p := fmt.Println

	count := make(chan int)
	wg.Add(2)
	// start the Goroutines
	p("Start Goroutines")
	go printCounts("One", count)
	go printCounts("TWO", count)

	// start the channel
	p("start the channel")
	count <- 1

	// closing sequence
	p("waiting to finish...")
	wg.Wait()
	p("Terminating program")

}
