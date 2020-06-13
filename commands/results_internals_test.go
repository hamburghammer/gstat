package commands

import (
	"errors"
	"testing"
)

func TestRmFirstLastBracket(t *testing.T) {
	t.Run("should remove first and last bracket", func(t *testing.T) {
		got := rmFirstLastBracket("{test}")
		want := "test"

		assertEqualString(got, want, t)
	})

	t.Run("should remove first and last bracket", func(t *testing.T) {
		got := rmFirstLastBracket("{test}{}")
		want := "test}{"

		assertEqualString(got, want, t)
	})
}

func TestCollectionEquals(t *testing.T) {

	t.Run("should be true if Results are equals", func(t *testing.T) {
		results := []string{"foo"}
		c1 := collection{Results: results}
		c2 := collection{Results: results}

		got := c1.collectionEquals(c2)
		want := true

		assertEqualBool(got, want, t)
	})

	t.Run("should be false if Results have different lengths", func(t *testing.T) {
		results1 := []string{"foo"}
		results2 := []string{"foo", "bar"}
		c1 := collection{Results: results1}
		c2 := collection{Results: results2}

		got := c1.collectionEquals(c2)
		want := false

		assertEqualBool(got, want, t)
	})

	t.Run("should be false if Results have different strings", func(t *testing.T) {
		results1 := []string{"foo"}
		results2 := []string{"bar"}
		c1 := collection{Results: results1}
		c2 := collection{Results: results2}

		got := c1.collectionEquals(c2)
		want := false

		assertEqualBool(got, want, t)
	})

	t.Run("should be true if Errs are equals", func(t *testing.T) {
		errs := []error{errors.New("foo")}
		c1 := collection{Errs: errs}
		c2 := collection{Errs: errs}

		got := c1.collectionEquals(c2)
		want := true

		assertEqualBool(got, want, t)
	})

	t.Run("should be false if Errs have different lengths", func(t *testing.T) {
		errs1 := []error{errors.New("foo")}
		errs2 := []error{errors.New("foo"), errors.New("bar")}
		c1 := collection{Errs: errs1}
		c2 := collection{Errs: errs2}

		got := c1.collectionEquals(c2)
		want := false

		assertEqualBool(got, want, t)
	})

	t.Run("should be false if Errs have different errors", func(t *testing.T) {
		errs1 := []error{errors.New("foo")}
		errs2 := []error{errors.New("bar")}
		c1 := collection{Errs: errs1}
		c2 := collection{Errs: errs2}

		got := c1.collectionEquals(c2)
		want := false

		assertEqualBool(got, want, t)
	})

	t.Run("should be false if Errs have different errors one being nil", func(t *testing.T) {
		errs1 := []error{errors.New("foo")}
		errs2 := []error{nil}
		c1 := collection{Errs: errs1}
		c2 := collection{Errs: errs2}

		got := c1.collectionEquals(c2)
		want := false

		assertEqualBool(got, want, t)
	})
}

func assertEqualString(got, want string, t *testing.T) {
	if got != want {
		t.Errorf("Want: '%s' but got: '%s'", want, got)
	}
}

func assertEqualBool(got, want bool, t *testing.T) {
	if got != want {
		t.Errorf("Want: '%t' but got: '%t'", want, got)
	}
}
