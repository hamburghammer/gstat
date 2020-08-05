package commands_test

import (
	"errors"
	"testing"

	"github.com/hamburghammer/gstat/args"
	"github.com/hamburghammer/gstat/commands"
	"github.com/stretchr/testify/assert"
)

func TestProcessExec(t *testing.T) {
	t.Run("should test for the Argument for Process", func(t *testing.T) {
		arguments := args.Arguments{Processes: false}
		got, err := commands.Processes{}.Exec(arguments)

		assert.Nil(t, err, "no error expected")
		assert.Equal(t, 0, len(got))
	})

	t.Run("should exit if the getting the process data returns an error", func(t *testing.T) {
		wantErr := errors.New("getting data error")
		arguments := args.Arguments{Processes: true}
		mockProcessData := func() ([]*commands.Process, error) { return nil, wantErr }
		_, err := commands.Processes{ReadProcesses: mockProcessData}.Exec(arguments)

		assert.NotNil(t, err, "an error was expected")
		assert.Equal(t, wantErr, err)
	})
}
