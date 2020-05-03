package main

import (
	"fmt"

	"github.com/hamburghammer/gstat/args"
)

func main() {
	fmt.Println("Hello, World!")
	args := args.Parse()
	fmt.Printf("Parsed args:\n	%+v\n", args)
}
