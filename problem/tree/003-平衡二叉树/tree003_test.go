package tree

import "testing"


func Test_isBalance(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalance(tt.args.root); got != tt.want {
				t.Errorf("isBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}
