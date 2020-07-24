package commands

import (
	"encoding/json"
	"sort"

	"github.com/hamburghammer/gstat/args"
	"github.com/shirou/gopsutil/process"
)

// Processes holds the function to get the process list
type Processes struct {
	ReadProcesses func() ([]*Process, error)
}

// NewProcesses is a factory ctor to build a Processes struct
func NewProcesses() Processes {
	return Processes{ReadProcesses: getProcesses}
}

// getProcesses maps the process.Process array to a local Process struct
func getProcesses() ([]*Process, error) {
	processes, err := process.Processes()

	p := make([]*Process, 0, len(processes))

	for _, process := range processes {
		p = append(p, &Process{Pid: process.Pid, CPUPercent: process.CPUPercent, Name: process.Name})
	}

	return p, err
}

// Exec is the implementation of the execution interface to be able to be used as a command
func (p Processes) Exec(args args.Arguments) ([]byte, error) {
	if !args.Processes {
		return []byte{}, nil
	}

	processes, err := p.ReadProcesses()
	if err != nil {
		return []byte{}, err
	}

	processesWithCPU := make([]cpuProcess, 0, len(processes))
	for _, process := range processes {
		cpuPercent, err := process.CPUPercent()
		if err != nil {
			return []byte{}, err
		}
		name, err := process.Name()
		if err != nil {
			return []byte{}, err
		}

		processesWithCPU = append(processesWithCPU, cpuProcess{Pid: process.Pid, CPU: cpuPercent, Name: name})
	}

	sort.Sort(byCPU(processesWithCPU))

	data := struct{ Processes []cpuProcess }{Processes: processesWithCPU[0:10]}
	return json.Marshal(data)

}

type cpuProcess struct {
	Name string
	Pid  int32
	CPU  float64
}

type byCPU []cpuProcess

func (c byCPU) Len() int           { return len(c) }
func (c byCPU) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c byCPU) Less(i, j int) bool { return c[i].CPU > c[j].CPU }

// Process is an adapter struct for the external process struct from github.com/shirou/gopsutil/process
type Process struct {
	Pid        int32
	Name       func() (string, error)
	CPUPercent func() (float64, error)
}
