package output

import (
	"github.com/hamburghammer/gstat/args"
	"github.com/hamburghammer/gstat/commands"
)

type Collection struct {
	Date    string                `json:"date,omitempty"`
	CPU     float64               `json:"cpu,omitempty"`
	Disk    commands.Memory       `json:"disk,omitempty"`
	Mem     commands.Memory       `json:"mem,omitempty"`
	Process []commands.CPUProcess `json:"process,omitempty"`
}

type Result struct {
	args args.Arguments
}

func (r Result) GatherResults() (Collection, []error) {
	var err error
	errs := make([]error, 0, 10)
	collection := Collection{}

	collection.Date = commands.NewDate().PureExec(r.args)
	collection.CPU, err = commands.NewCPU().PureExec(r.args)
	if err != nil {
		errs = append(errs, err)
	}
	collection.Disk, err = commands.NewDisk().PureExec(r.args)
	if err != nil {
		errs = append(errs, err)
	}
	collection.Mem, err = commands.NewMem().PureExec(r.args)
	if err != nil {
		errs = append(errs, err)
	}
	collection.Process, err = commands.NewProcesses().PureExec(r.args)
	if err != nil {
		errs = append(errs, err)
	}

	return collection, errs
}

func NewResult(args args.Arguments) Result {
	return Result{args: args}
}
