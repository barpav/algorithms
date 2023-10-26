package problem941

import "testing"

func Test_validMountainArray(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "[2,1]",
			arr:  []int{2, 1},
			want: false,
		},
		{
			name: "[3,5,5]",
			arr:  []int{3, 5, 5},
			want: false,
		},
		{
			name: "[0,3,2,1]",
			arr:  []int{0, 3, 2, 1},
			want: true,
		},
		{
			name: "[0,2,3,4,5,2,1,0]",
			arr:  []int{0, 2, 3, 4, 5, 2, 1, 0},
			want: true,
		},
		{
			name: "[3,2,1]",
			arr:  []int{3, 2, 1},
			want: false,
		},
		{
			name: "[1,2,3]",
			arr:  []int{1, 2, 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.arr); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
