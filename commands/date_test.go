package commands_test

import (
	"testing"

	"github.com/hamburghammer/gstat/args"
	"github.com/hamburghammer/gstat/commands"
	"github.com/stretchr/testify/assert"
)

func TestDate(t *testing.T) {
	t.Run("", func(t *testing.T) {
		customTime := func() string { return "2020-08-09T17:43:31+02:00" }
		date := commands.Date{GetTime: customTime}

		got, err := date.Exec(args.Arguments{})
		want := "{\"Date\":\"2020-08-09T17:43:31+02:00\"}"

		assert.Nil(t, err)
		assert.Equal(t, want, string(got))
	})
}
