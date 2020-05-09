package proc

import (
	"fmt"
	"time"
)

// TotalCPU returns the first entry from the return array form the given function
func TotalCPU(c chan float64, cpuPercent func(interval time.Duration, percpu bool) ([]float64, error)) {
	total, err := cpuPercent(time.Millisecond*500, false)

	if err != nil {
		fmt.Printf("Something went wrong reading the cpu: %v \n", err)
	}

	fmt.Printf("CPU: %v\n", total)

	c <- total[0]
}
