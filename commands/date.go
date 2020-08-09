package commands

import (
	"encoding/json"
	"time"

	"github.com/hamburghammer/gstat/args"
)

// Date is a the configuration struct to execute the date command
type Date struct {
	// GetTime is the function to get the actual Time in string format.
	// It should be use to replace/cusomise the time output.
	GetTime func() string
}

// NewDate is a convinice constructor for the Date struct.
// It sets the GetTime function to standard formatting.
func NewDate() Date {
	return Date{GetTime: getFormattedTime}
}

// Exec is the implementation of the execution interface for the Date struct.
func (d Date) Exec(args args.Arguments) ([]byte, error) {
	data := struct{ Date string }{Date: d.GetTime()}
	return json.Marshal(data)
}

func getFormattedTime() string {
	return time.Now().Format(time.RFC3339)
}
