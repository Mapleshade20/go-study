package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// message, err := greetings.Hello("Maple")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(message)
	fmt.Println(quote.Go())

	names := []string{"Gladys", "Samantha", "Maple"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
