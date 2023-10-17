package problem27

import (
	"reflect"
	"sort"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		k    int
		want []int
	}{
		{
			name: "[3,2,2,3], 3",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			k:    2,
			want: []int{2, 2},
		},
		{
			name: "[0,1,2,2,3,0,4,2], 2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			k:    5,
			want: []int{0, 0, 1, 3, 4},
		},
	}

	var k int
	var got []int
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k = removeElement(tt.args.nums, tt.args.val)

			if k != tt.k {
				t.Errorf("removeElement() = %v, want %v", k, tt.k)
			}

			got = tt.args.nums[:k]
			sort.Ints(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
