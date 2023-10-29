package problem905

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParity(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "[3,1,2,4]",
			nums: []int{3, 1, 2, 4},
			want: []int{2, 4, 3, 1},
		},
		{
			name: "[2, 4, 3, 1]",
			nums: []int{2, 4, 3, 1},
			want: []int{2, 4, 3, 1},
		},
		{
			name: "[2, 4, 3, 1, 0]",
			nums: []int{2, 4, 3, 1, 0},
			want: []int{2, 4, 0, 1, 3},
		},
		{
			name: "[0]",
			nums: []int{0},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParity(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParity() = %v, want %v", got, tt.want)
			}
		})
	}
}
