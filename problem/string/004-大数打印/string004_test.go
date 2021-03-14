package string

import (
	"reflect"
	"testing"
)

func Test_printNum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test1", args{num: 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printNum(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
