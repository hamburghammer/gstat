package main

import (
	"fmt"

	"github.com/hamburghammer/gstat/args"
	"github.com/hamburghammer/gstat/commands"
)

func main() {
	args := args.Parse()

	result := commands.NewResult(args)
	executs := []commands.Executor{commands.NewCPU()}
	output := result.ExecCommands(executs)
	fmt.Println(output.Collection.Results)
}
