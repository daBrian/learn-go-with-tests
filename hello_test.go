package main

import "testing"

func TestName(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, bring string, want string) {
		t.Helper()
		if got := Hello(bring, ""); got != want {
			t.Errorf("Hello() = %v, want %v", got, want)
		}
	}

	tests := []struct {
		name  string
		bring string
		want  string
	}{
		{"Greet world", "world", "Hello, world"},
		{"Greet world without parameter", "", "Hello, world"},
		{"Greet Chris", "Chris", "Hello, Chris"},
		{"Greet Brian", "Brian", "Hello, Brian"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertCorrectMessage(t, tt.bring, tt.want)
		})
	}
}

func TestLanguage(t *testing.T) {
	type args struct {
		name     string
		language string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "in Spanish",
			args: args{name: "Elodie", language: "Spanish"},
			want: "Hola, Elodie",
		},

		{
			name: "in French",
			args: args{name: "Elodie", language: "French"},
			want: "Bonjour, Elodie",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.name, tt.args.language); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
