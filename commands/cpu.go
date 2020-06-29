package commands

import (
	"encoding/json"
	"time"

	"github.com/hamburghammer/gstat/args"
	"github.com/hamburghammer/gstat/errors"
	"github.com/shirou/gopsutil/cpu"
)

// OperationKeyCPUReading represents the key for the Operation field of an CPUReadingError
const OperationKeyCPUReading = "CPUReading"

// CPU holds the config to get the cpu load in percentage
type CPU struct {
	TimeInMilSec int
	ReadCPUStat  func(interval time.Duration, percpu bool) ([]float64, error)
}

// NewCPU creates a new cpu percentage struct
func NewCPU() CPU {
	return CPU{TimeInMilSec: 500, ReadCPUStat: cpu.Percent}
}

// Exec gets the cpu value and maps it to the executiondata struct
func (c CPU) Exec(args args.Arguments) ([]byte, error) {
	if !args.CPU {
		return []byte{}, nil
	}
	total, err := c.ReadCPUStat(time.Millisecond*time.Duration(c.TimeInMilSec), false)
	if err != nil {
		return []byte{}, errors.BaseError{OperationKeyCPUReading, err.Error()}
	}
	data := struct{ CPU float64 }{CPU: total[0]}
	return json.Marshal(data)
}
