package gocode

import "testing"

func Test_printNumberAndLetter(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printNumberAndLetter()
		})
	}
}