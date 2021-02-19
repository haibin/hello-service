package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("hello world")
	if err := run(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func run() error {
	return nil
}