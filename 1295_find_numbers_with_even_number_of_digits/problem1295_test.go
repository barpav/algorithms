package problem1295

import "testing"

func Test_findNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "[12,345,2,6,7896]",
			nums: []int{12, 345, 2, 6, 7896},
			want: 2,
		},
		{
			name: "[555,901,482,1771]",
			nums: []int{555, 901, 482, 1771},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumbers(tt.nums); got != tt.want {
				t.Errorf("findNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
