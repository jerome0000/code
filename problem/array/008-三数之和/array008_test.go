package array

import (
	"reflect"
	"testing"
)

func Test_thirdSum(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{
			arr:    []int{1, 2, 3, 4, 5, 6},
			target: 6,
		}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thirdSum(tt.args.arr, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("thirdSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
