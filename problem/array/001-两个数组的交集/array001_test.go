package array

import (
	"reflect"
	"testing"
)

func Test_intersectV1(t *testing.T) {
	type args struct {
		num1 []int
		num2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{
			num1: []int{1, 2, 3, 4},
			num2: []int{2, 3, 4, 5, 6},
		}, []int{2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersectV1(tt.args.num1, tt.args.num2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intersectV2(t *testing.T) {
	type args struct {
		num1 []int
		num2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{
			num1: []int{1, 2, 3, 4},
			num2: []int{2, 3, 4, 5, 6},
		}, []int{2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersectV2(tt.args.num1, tt.args.num2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
