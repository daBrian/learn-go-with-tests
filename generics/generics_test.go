package generics_test

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
}
func TestAssertFunctionsWithStrings(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, "mary", "mary")
		AssertNotEqual(t, "mary", "john")
	})
}
func AssertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func AssertNotEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %+v", got)
	}
}
