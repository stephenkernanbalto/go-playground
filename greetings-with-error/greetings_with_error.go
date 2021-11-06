package greetings_with_error

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		// we expect a response list, so we return (empty string, error)
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v", name)
	// we expect a response list, so we return (message, nil - meaning no error)
	return message, nil
}