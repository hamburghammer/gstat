package main

import (
	"fmt"

	"github.com/hamburghammer/gstat/proc"
)

func main() {
	fmt.Println("Hello, World!")

	total, err := proc.TotalCPU()

	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("from the channel: %f\n", total)
}
