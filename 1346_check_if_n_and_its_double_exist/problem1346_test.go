package problem1346

import "testing"

func Test_checkIfExist(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "[10,2,5,3]",
			arr:  []int{10, 2, 5, 3},
			want: true,
		},
		{
			name: "[5,2,10,3]",
			arr:  []int{5, 2, 10, 3},
			want: true,
		},
		{
			name: "[3,1,7,11]",
			arr:  []int{3, 1, 7, 11},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfExist(tt.arr); got != tt.want {
				t.Errorf("checkIfExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
