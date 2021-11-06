package main

import (
	"fmt"
	"log"

	"go_playground/greetings_multiple"
	"go_playground/hello/local_package"
)

func main() {
	log.SetPrefix("greetings_with_slices: ")
	log.SetFlags(0)

    names := []string{"Stephen", "Austin"}
	message, err := greetings_multiple.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	local_package.Print("this is a message")
}