package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hamburghammer/gstat/args"
	"github.com/hamburghammer/gstat/commands"
	"github.com/hamburghammer/gstat/output"
)

func main() {
	args := args.Parse()

	output, err := newOutput(args)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(output))
}

func oldOutput(args args.Arguments) string {
	result := commands.NewResult(args)
	executs := []commands.Executor{
		commands.NewDate(),
		commands.NewCPU(),
		commands.NewMem(),
		commands.NewDisk(),
		commands.NewProcesses(),
	}
	output := result.ExecCommands(executs)

	return formatToJSON(output.Collection.Results)
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

func newOutput(args args.Arguments) ([]byte, error) {
	result := output.NewResult(args)
	col, errs := result.GatherResults()
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		return []byte{}, fmt.Errorf("Collection the stats produced an error")
	}

	return json.Marshal(col)
}
