package main

import (
    "fmt"
    "example.com/greetings"
	"log"
)

func main() {

	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Lewis")

	if err != nil {
		log.Fatal(err)
	}
    // Get a greeting message and print it.
    // message := greetings.Hello("Gladys")
    fmt.Println(message)
}