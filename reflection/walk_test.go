package reflection

import (
	"reflect"
	"testing"
)

func Test_walk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		}, {
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		}, {
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		}, {
			"Nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		}, {
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
	}
	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			var got []string
			walk(tt.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, tt.ExpectedCalls) {
				t.Errorf("got %v want %v", got, tt.ExpectedCalls)
			}
		})
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
