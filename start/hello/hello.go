package main

import (
	"fmt"
	"start/greetings"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
