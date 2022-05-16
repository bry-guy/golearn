package main

import (
	"fmt"
	"log"
	"strings"

	"bry-guy.net/greetings"
)

func main() {
	// Disable printing time, source file, and line number.
	log.SetFlags(0)
	log.SetPrefix(strings.ToUpper("hello: "))

	names := []string{"GladOS", "Foo"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range messages {
		fmt.Println(v)
	}
}
