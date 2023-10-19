package problem26

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "[1,1,2]",
			nums: []int{1, 1, 2},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "[0,0,1,1,1,2,2,3,3,4]",
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			k:    5,
			want: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := removeDuplicates(tt.nums)

			if k != tt.k {
				t.Errorf("k = %d, want %d", k, tt.k)
			}

			if got := tt.nums[:k]; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
