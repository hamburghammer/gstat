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

// Exec is the implementation of the execution interface to be able to be used as a command
func (p Processes) Exec(args args.Arguments) ([]byte, error) {
	processes, err := p.PureExec(args)
	if err != nil {
		return []byte{}, err
	}

	data := struct{ Processes []CPUProcess }{Processes: processes}
	return json.Marshal(data)
}

func (p Processes) PureExec(args args.Arguments) ([]CPUProcess, error) {
	if !args.Processes {
		return []CPUProcess{}, nil
	}

	processes, err := p.ReadProcesses()
	if err != nil {
		return []CPUProcess{}, err
	}

	processesWithCPU, err := getProcessesCPUInfos(processes)
	if err != nil {
		return []CPUProcess{}, err
	}

	sort.Sort(byCPU(processesWithCPU))

	return getFirstTenOrLess(processesWithCPU), nil
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

func getFirstTenOrLess(array []CPUProcess) []CPUProcess {
	if len(array) >= 9 {
		return array[0:10]
	}
	return array
}

func getProcessesCPUInfos(processes []*Process) ([]CPUProcess, error) {
	processesWithCPU := make([]CPUProcess, 0, len(processes))
	for _, process := range processes {
		processCPUInfo, err := getProcessCPUInfos(process)
		if err != nil {
			return processesWithCPU, err
		}
		processesWithCPU = append(processesWithCPU, *processCPUInfo)
	}
	return processesWithCPU, nil
}

func getProcessCPUInfos(process *Process) (*CPUProcess, error) {
	cpuPercent, err := process.CPUPercent()
	if err != nil {
		return &CPUProcess{}, err
	}
	name, err := process.Name()
	if err != nil {
		return &CPUProcess{}, err
	}
	return &CPUProcess{Pid: process.Pid, CPU: cpuPercent, Name: name}, nil
}

type CPUProcess struct {
	Name string
	Pid  int32
	CPU  float64
}

type byCPU []CPUProcess

func (c byCPU) Len() int           { return len(c) }
func (c byCPU) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c byCPU) Less(i, j int) bool { return c[i].CPU > c[j].CPU }

// Process is an adapter struct for the external process struct from github.com/shirou/gopsutil/process
type Process struct {
	Pid        int32
	Name       func() (string, error)
	CPUPercent func() (float64, error)
}
