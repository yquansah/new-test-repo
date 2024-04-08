package main

import (
	"fmt"
	"log"
)

func run() error {
	fmt
	fmt.Println("HELLO WORLD")
	return nil
}

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
