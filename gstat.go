package main

import (
	"github.com/hamburghammer/gstat/args"
	"github.com/hamburghammer/gstat/commands"
)

func main() {
	args := args.Parse()

	result := commands.NewResult(args)
	result.ExecCommands()
}
