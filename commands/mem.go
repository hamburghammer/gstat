package commands

import (
	"encoding/json"

	"github.com/hamburghammer/gstat/args"
	"github.com/shirou/gopsutil/mem"
)

// Memory usage representation.
type Memory struct {
	Used  uint64 `json:"used"`
	Total uint64 `json:"total"`
}

// Mem holds the memory usage for the json transformation
type Mem struct {
	ReadVirtualMemoryStat func() (*mem.VirtualMemoryStat, error)
}

// NewMem is a constructor for the Mem struct
func NewMem() Mem {
	return Mem{mem.VirtualMemory}
}

// Exec gets the mem value and maps it to the executiondata struct
func (m Mem) Exec(args args.Arguments) ([]byte, error) {
	if !args.Mem {
		return []byte{}, nil
	}

	mem, err := m.ReadVirtualMemoryStat()
	if err != nil {
		return []byte{}, err
	}

	usage := Memory{Used: bytesToMegaByte(mem.Used), Total: bytesToMegaByte(mem.Total)}
	data := struct {
		Mem Memory `json:"mem"`
	}{usage}
	return json.Marshal(data)
}

func bytesToMegaByte(bytes uint64) uint64 {
	kb := bytes / 1024
	mb := kb / 1024
	return mb
}
