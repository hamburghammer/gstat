package commands

import (
	"strings"

	"github.com/hamburghammer/gstat/args"
)

// Result gethers all the results for all commands
type Result struct {
	args.Arguments
	Collection collection
}

type collection struct {
	Results []string
	Errs    []error
}

func (c collection) collectionEquals(otherCollection collection) bool {
	if len(c.Results) != len(otherCollection.Results) {
		return false
	}
	for i, s := range c.Results {
		if s != otherCollection.Results[i] {
			return false
		}
	}

	if len(c.Errs) != len(otherCollection.Errs) {
		return false
	}
	for i, s := range c.Errs {
		if s != otherCollection.Errs[i] {
			return false
		}
	}

	return true
}

// NewResult creates new result struct
func NewResult(a args.Arguments) Result {
	return Result{Arguments: a}
}

// ExecCommands runs all commands
func (r Result) ExecCommands(executors []Executor) Result {
	results := make([]string, 0, len(executors))
	errors := make([]error, 0, len(executors))

	for _, executor := range executors {
		output, err := executor.Exec(r.Arguments)
		s := string(output)
		if err != nil {
			errors = append(errors, err)
		} else {
			results = append(results, rmFirstLastBracket(s))
		}
	}

	r.Collection.Errs = errors
	r.Collection.Results = results
	return r
}

// ResultEquals checks for field equality
func (r Result) ResultEquals(otherResult Result) bool {
	if !r.Arguments.Equals(otherResult.Arguments) {
		return false
	}
	if !r.Collection.collectionEquals(otherResult.Collection) {
		return false
	}
	return true
}

func rmFirstLastBracket(s string) string {
	s = strings.Replace(s, "{", "", 1)
	s = reverse(strings.Replace(reverse(s), "}", "", 1))
	return s
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
