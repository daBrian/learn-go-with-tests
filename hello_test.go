package main

import "testing"

func TestHello(t *testing.T) {
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
			if got := Hello(tt.bring); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
