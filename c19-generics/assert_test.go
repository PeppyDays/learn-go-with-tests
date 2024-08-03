package c19_test

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "Arine")
	})
}

func AssertEqual[T comparable](t *testing.T, expected T, actual T) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}

func AssertNotEqual[T comparable](t *testing.T, expected T, actual T) {
	t.Helper()
	if expected == actual {
		t.Errorf("not expected %v, but got %v", expected, actual)
	}
}
