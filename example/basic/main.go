package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sweetbbak/go-nude"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		println("usage: basic <path/to/image>")
	}

	isNude, err := nude.IsNude(args[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("isNude = %v\n", isNude)
}
