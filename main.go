package main

import (
	"fmt"

	"github.com/hamburghammer/gstat/args"
	"github.com/hamburghammer/gstat/commands"
)

func main() {
	args := args.Parse()

	result := commands.NewResult(args)
	executs := []commands.Executor{commands.NewCPU(), commands.NewMem(), commands.NewDisk()}
	output := result.ExecCommands(executs)

	fmt.Println(formatToJSON(output.Collection.Results))
}

func formatToJSON(strings []string) string {
	stringBuilder := "{"
	elements := len(strings)

	for i, s := range strings {
		stringBuilder = stringBuilder + s
		if i != (elements - 1) {
			stringBuilder = stringBuilder + ","
		}
	}
	stringBuilder = stringBuilder + "}"

	return stringBuilder
}
