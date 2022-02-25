package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, bring string, want string) {
		t.Helper()
		if got := Hello(bring); got != want {
			t.Errorf("Hello() = %v, want %v", got, want)
		}
	}

	tests := []struct {
		name  string
		bring string
		want  string
	}{
		{"Greet world", "world", "Hello, world"},
		{"Greet Chris", "Chris", "Hello, Chris"},
		{"Greet Brian", "Brian", "Hello, Brian"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertCorrectMessage(t, tt.bring, tt.want)
		})
	}
}
