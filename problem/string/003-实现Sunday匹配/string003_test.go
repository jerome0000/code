package string

import "testing"

func Test_str(t *testing.T) {
	type args struct {
		s string
		a string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{s: "asthisistesta", a: "as"}, 0},
		{"test1", args{s: "aasthisistest", a: "as"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := str(tt.args.s, tt.args.a); got != tt.want {
				t.Errorf("str() = %v, want %v", got, tt.want)
			}
		})
	}
}
