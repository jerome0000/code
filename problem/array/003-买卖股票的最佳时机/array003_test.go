package array

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		price []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{price: []int{7, 1, 5, 3, 6, 4}}, 5},
		{"test2", args{price: []int{1, 2, 3, 4, 5}}, 4},
		{"test3", args{price: []int{7, 6, 4, 3, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.price); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
