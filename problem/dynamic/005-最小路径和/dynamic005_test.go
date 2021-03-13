package dynamic

import "testing"

func Test_minPath(t *testing.T) {
	type args struct {
		arr [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{arr: [][]int{
			{1, 2, 4, 5},
			{1, 2, 3, 4},
		}}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPath(tt.args.arr); got != tt.want {
				t.Errorf("minPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
