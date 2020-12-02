# Unbuffered channels in Go

An unbuffered channel provides synchronous communication among goroutines, which ensures message deilivery among them. Message sending is only permitted if there is a receiver waiting to receive the message.

In the case we see here for unbuffered channels the sync waits for the group to be ready to receive. We see 2 passed into the waitGroup and the unbuffered channel is passed between the two goroutines. The most important aspect of an unbuffered channel is that it is blocking.

```go
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

```

# Buffered channels

A buffered channel can have values sent up to its capacity without having a received to receive them. See the trivial bellow option where we have created a channel with the capacity of 2. We can then hold two values beong passed in to it before we call any receiver afterwards. The difference is that this capcity is specified and can be used without a receiver up until the capcity of the channel. An unbuffered chnnel may be thought of as having a capcity of 1 and therefore it requires what is sent in to be received before it can be reused to send more.

```go
package main

import "fmt"

func main() {
	messages := make(chan string, 5)

	messages <- "Golang"
	messages <- "for the win"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
```
