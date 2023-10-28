package problem283

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "[0,1,0,3,12]",
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "[0,1,0,0,3,12]",
			nums: []int{0, 1, 0, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0, 0},
		},
		{
			name: "[1,1,0,0,3,12]",
			nums: []int{1, 1, 0, 0, 3, 12},
			want: []int{1, 1, 3, 12, 0, 0},
		},
		{
			name: "[1,1,0,4,6,0,3,12,0,0]",
			nums: []int{1, 1, 0, 4, 6, 0, 3, 12, 0, 0},
			want: []int{1, 1, 4, 6, 3, 12, 0, 0, 0, 0},
		},
		{
			name: "[0]",
			nums: []int{0},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("got = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
