package problem9

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{
			name: "121",
			x:    121,
			want: true,
		},
		{
			name: "-121",
			x:    -121,
			want: false,
		},
		{
			name: "10",
			x:    10,
			want: false,
		},
		{
			name: "11",
			x:    11,
			want: true,
		},
		{
			name: "1",
			x:    1,
			want: true,
		},
		{
			name: "1221",
			x:    1221,
			want: true,
		},
		{
			name: "1231",
			x:    1231,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
