package problem2

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
		name string
		args args
		want *ListNode
	}{
		{
			name: "l1 = [2,4,3], l2 = [5,6,4]",
			args: args{
				l1: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 2,
						},
					},
				},
				l2: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 8,
					},
				},
			},
		},
		{
			name: "l1 = [0], l2 = [0]",
			args: args{
				l1: &ListNode{Val: 0},
				l2: &ListNode{Val: 0},
			},
			want: &ListNode{Val: 0},
		},
		{
			name: "l1 = [9], l2 = [9,0,1]",
			args: args{
				l1: &ListNode{
					Val: 9,
				},
				l2: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
			want: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
