package problem13

import "testing"

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "III",
			s:    "III",
			want: 3,
		},
		{
			name: "IX",
			s:    "IX",
			want: 9,
		},
		{
			name: "X",
			s:    "X",
			want: 10,
		},
		{
			name: "XLIV",
			s:    "XLIV",
			want: 44,
		},
		{
			name: "LVIII",
			s:    "LVIII",
			want: 58,
		},
		{
			name: "DCCLXXXIX",
			s:    "DCCLXXXIX",
			want: 789,
		},
		{
			name: "MCMXCIV",
			s:    "MCMXCIV",
			want: 1994,
		},
		{
			name: "MMMCMXCIX",
			s:    "MMMCMXCIX",
			want: 3999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
