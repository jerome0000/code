package linked_list004

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name     string
		args     args
		wantHead *ListNode
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHead := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(gotHead, tt.wantHead) {
				t.Errorf("addTwoNumbers() = %v, want %v", gotHead, tt.wantHead)
			}
		})
	}
}
