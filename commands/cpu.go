package commands

import (
	"encoding/json"
	"fmt"
	"time"

	e "github.com/hamburghammer/gstat/errors"
	"github.com/shirou/gopsutil/cpu"
)

// OperationKeyCPUReading represents the key for the Operation field of an CPUReadingError
const OperationKeyCPUReading = "CPUReading"

// CPUPercentage holds the information to build the output json
type CPUPercentage struct {
	CPU float64
}

// NewCPU creates a new cpu percentage struct
func NewCPU() CPUPercentage {
	return CPUPercentage{}
}

// Exec gets the cpu value and maps it to the executiondata struct
func (c CPUPercentage) Exec() ([]byte, error) {
	total, err := TotalCPU()
	if err != nil {
		return []byte{}, err
	}
	c.CPU = total
	fmt.Println(c)
	return json.Marshal(c)
}

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
