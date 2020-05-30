package proc

import (
	"fmt"
	"time"

	e "github.com/hamburghammer/gstat/errors"
	"github.com/shirou/gopsutil/cpu"
)

// OperationKeyCPUReading represents the key for the Operation field of an CPUReadingError
const OperationKeyCPUReading = "CPUReading"

// TotalCPU returns the first entry of the return array form the given function
func TotalCPU() (float64, error) {
	total, err := cpu.Percent(time.Millisecond*500, false)

	if err != nil {
		wrappedError := fmt.Errorf("Something went wrong reading the cpu: %w", err)
		return float64(0), e.BaseError{Operation: OperationKeyCPUReading, Message: wrappedError.Error()}
	}

	if len(total) != 1 {
		return float64(0), e.BaseError{Operation: OperationKeyCPUReading, Message: "No CPU data was found. Please check the HOST_PROC env to point to the right directory."}
	}

	return total[0], nil
}
