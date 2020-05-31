package commands

import "github.com/hamburghammer/gstat/args"

// Executor is a functional interface to execute an command and return the result as a json string.
type Executor interface {
	// Exec executes something and returns the result as a byte array of json and an error if something unexpected happened.
	Exec(args.Arguments) ([]byte, error)
}
