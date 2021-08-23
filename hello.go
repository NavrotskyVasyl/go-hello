package main

import (
	"fmt"
	"log"

	"github.com/NavrotskyVasyl/go-greetings"
)


func main() {
	log.SetPrefix("greetings: ")
    log.SetFlags(0)

    message, err := greetings.Hello("Unit")
    if err != nil {
    	log.Fatal(err)
    }

    fmt.Println(message)
}