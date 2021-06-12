package sort001

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{arr: []int{1, 2, 4, 5, 6, 7, 2, 1, 3, 4, 6}}, []int{1, 1, 2, 2, 3, 4, 4, 5, 6, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
