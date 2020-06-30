package main

import "testing"

func TestFormatToJSON(t *testing.T) {
	strings := []string{"foo", "bar"}

	got := formatToJSON(strings)
	want := "{foo,bar}"

	if got != want {
		t.Errorf("Want: '%s' but got: '%s'", want, got)
	}
}
