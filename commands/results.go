package commands

import (
	"fmt"
	"strings"

	"github.com/hamburghammer/gstat/args"
)

// Result gethers all the results for all commands
type Result struct {
	args.Arguments
	collection
}

type collection struct {
	results []string
	errs    []error
}

// NewResult creates new result struct
func NewResult(a args.Arguments) Result {
	return Result{Arguments: a}
}

// ExecCommands runs all commands
func (r *Result) ExecCommands() {
	cpu, err := NewCPU().Exec(r.Arguments)
	if err != nil {
		panic(err)
	}
	cpustr := rmFirstLastBracket(string(cpu))

	fmt.Printf("{%s}\n", cpustr)
}

func rmFirstLastBracket(s string) string {
	s = strings.Replace(s, "{", "", 1)
	return strings.Replace(s, "}", "", 1)
}
