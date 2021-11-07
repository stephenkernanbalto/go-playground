package main

import (
	"fmt"
	"time"
)

func main() {
	// sending in channels is blocking until receiving happens.
	// we can add a second param as a buffer_count to allow room for continued channel operation.
	// then, a channel can send data n times where n <= buffer_count
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"

	// receives a message from channel
	msg := <- c
	fmt.Println(msg)

	// receives a second message from channel
	msg = <- c
	fmt.Println(msg)
}