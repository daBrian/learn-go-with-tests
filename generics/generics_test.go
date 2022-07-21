package generics_test

import (
	. "github.com/daBrian/learn-go-with-tests/generics"
	"testing"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		MAssertEqual(t, 1, 1)
		MAssertNotEqual(t, 1, 2)
	})
	t.Run("asserting on strings", func(t *testing.T) {
		MAssertEqual(t, "mary", "mary")
		MAssertNotEqual(t, "mary", "john")
	})

	//t.Run("assert mixed", func(t *testing.T) {
	//	MAssertEqual(t, 1, "1")
	//})
}
