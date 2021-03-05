package array

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test", args{arr: []string{}}, ""},
		{"test", args{arr: []string{"a", "aa111"}}, "a"},
		{"test", args{arr: []string{"aaa", "aa111"}}, "aa"},
		{"test", args{arr: []string{"aaa", "aaa111"}}, "aaa"},
		{"test", args{arr: []string{"aaa", "aaa111", "aaa222"}}, "aaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.arr); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
