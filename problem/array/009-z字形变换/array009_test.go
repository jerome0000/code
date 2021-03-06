package array

import "testing"

func Test_z(t *testing.T) {
	type args struct {
		s    string
		rows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test", args{
			s:    "LEETCODEISHIRING",
			rows: 4,
		}, "LDREOEIIECIHNTSG"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := z(tt.args.s, tt.args.rows); got != tt.want {
				t.Errorf("z() = %v, want %v", got, tt.want)
			}
		})
	}
}
