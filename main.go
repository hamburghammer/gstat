package main

import (
	"fmt"

	"github.com/hamburghammer/gstat/proc"
	"github.com/shirou/gopsutil/cpu"
)

func main() {
	fmt.Println("Hello, World!")

	totalCPUCannel := make(chan float64)
	go proc.TotalCPU(totalCPUCannel, cpu.Percent)

	totalCPU := <-totalCPUCannel
	fmt.Printf("from the channel: %f\n", totalCPU)
}
