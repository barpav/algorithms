package problem1089

import (
	"reflect"
	"testing"
)

func Test_duplicateZeros(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "[1,0,2,3,0,4,5,0]",
			arr:  []int{1, 0, 2, 3, 0, 4, 5, 0},
			want: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			name: "[1,2,3]",
			arr:  []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicateZeros(tt.arr)
			if !reflect.DeepEqual(tt.arr, tt.want) {
				t.Errorf("arr = %v, want %v", tt.arr, tt.want)
			}
		})
	}
}
