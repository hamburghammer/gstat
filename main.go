package main

import (
	"fmt"

	"github.com/hamburghammer/gstat/commands"
)

func main() {
	fmt.Println("Hello, World!")

	total, err := commands.TotalCPU()

	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("from the channel: %f\n", total)
}
