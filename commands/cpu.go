package commands

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hamburghammer/gstat/args"
	e "github.com/hamburghammer/gstat/errors"
	"github.com/shirou/gopsutil/cpu"
)

// OperationKeyCPUReading represents the key for the Operation field of an CPUReadingError
const OperationKeyCPUReading = "CPUReading"

// CPU holds the config to get the cpu load in percentage
type CPU struct {
	TimeInMilSec int
}

// NewCPU creates a new cpu percentage struct
func NewCPU() CPU {
	return CPU{500}
}

// Exec gets the cpu value and maps it to the executiondata struct
func (c CPU) Exec(args args.Arguments) ([]byte, error) {
	if !args.CPU {
		return []byte{}, nil
	}
	total, err := c.TotalCPU()
	if err != nil {
		return []byte{}, err
	}
	data := struct{ CPU float64 }{CPU: total}
	return json.Marshal(data)
}

// TotalCPU returns the first entry of the return array form the given function
func (c CPU) TotalCPU() (float64, error) {
	total, err := cpu.Percent(time.Millisecond*time.Duration(c.TimeInMilSec), false)

	if err != nil {
		wrappedError := fmt.Errorf("Something went wrong reading the CPU: %w", err)
		return float64(0), e.BaseError{Operation: OperationKeyCPUReading, Message: wrappedError.Error()}
	}

	if len(total) != 1 {
		return float64(0), e.BaseError{
			Operation: OperationKeyCPUReading,
			Message:   "No CPU data was found. Please check the HOST_PROC env to point to the right directory."}
	}
	return total[0], nil
}
