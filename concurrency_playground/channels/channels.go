package channels

import (
	"fmt"
	"time"
)

func Solution() {
	c := make(chan string)
	go count("sheep", c)
	// receives a message from channel
	for msg := range c {
		fmt.Println(msg)
	}
}

func SolutionVerbose() {
	c := make(chan string)
	go count("sheep", c)
	// under the hood of lines 12-13, Go does this
	for {
		msg, open := <- c

		if !open {
			break
		}

		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i < 50; i++ {
		// send a message to channel
		c <- thing
		time.Sleep(time.Millisecond * 5)
	}

	close(c)
}