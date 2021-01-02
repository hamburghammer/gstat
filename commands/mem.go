package commands

import (
	"encoding/json"

	"github.com/hamburghammer/gstat/args"
	"github.com/shirou/gopsutil/mem"
)

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
	usage, err := m.PureExec(args)
	if err != nil {
		return []byte{}, err
	}

	data := struct {
		Mem Memory `json:"mem"`
	}{usage}
	return json.Marshal(data)
}

func (m Mem) PureExec(args args.Arguments) (Memory, error) {
	if !args.Mem {
		return Memory{}, nil
	}

	mem, err := m.ReadVirtualMemoryStat()
	if err != nil {
		return Memory{}, err
	}

	memory := Memory{Used: bytesToMegaByte(mem.Used), Total: bytesToMegaByte(mem.Total)}
	return memory, nil
}

func bytesToMegaByte(bytes uint64) uint64 {
	kb := bytes / 1024
	mb := kb / 1024
	return mb
}
