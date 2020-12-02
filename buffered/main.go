package main

import "fmt"

func main() {
	messages := make(chan string, 5)

	messages <- "Golang"
	messages <- "for the win"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
