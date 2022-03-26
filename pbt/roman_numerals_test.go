package pbt

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Arabic int
		Want   string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
	}
	for _, test := range cases {
		description := fmt.Sprintf("%d gets converted to %s", test.Arabic, test.Want)
		t.Run(description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)

			if got != test.Want {
				t.Errorf("got %q want %q", got, test.Want)
			}
		})
	}
}
