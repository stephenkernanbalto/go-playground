package select_statements

import (
	"fmt"
	"time"
)

func Problem() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	//
	go func() {
		for {
			c2 <- "Every 2s"
			time.Sleep(time.Second * 2)
		}
	}()

	// Sending and receiving are blocking operations for the channel.
	// In this case, c2 blocks for 2 seconds before printing every time.
	// So, our goroutines will always print c1, then c2, in a 2 second loop.
	for {
		fmt.Println(<- c1)
		fmt.Println(<- c2)
	}
}

func Solution() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2s"
			time.Sleep(time.Second * 2)
		}
	}()


	// We get around this with a select statement.
	// Select says to always receive from the channel that is ready to send.
	// In this case, it will receive from c1 every time c1 is ready to send a message.
	// Then, when c2 is ready, it will receive from c2 and go back to c1.
	for {
		select {
			case msg1 := <- c1:
				fmt.Println(msg1)
			case msg2 := <- c2:
				fmt.Println(msg2)
		}
	}
}