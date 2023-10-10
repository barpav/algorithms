package problem3

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "abcdef",
			s:    "abcdef",
			want: 6,
		},
		{
			name: "abcabcbb",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "bbbbb",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "pwwkew",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "dvdf",
			s:    "dvdf",
			want: 3,
		},
		{
			name: "anviaj",
			s:    "anviaj",
			want: 5,
		},
		{
			name: "ckilbkd",
			s:    "ckilbkd",
			want: 5,
		},
		{
			name: "cdd",
			s:    "cdd",
			want: 2,
		},
		{
			name: "abba",
			s:    "abba",
			want: 2,
		},
		{
			name: "tmmzuxt",
			s:    "tmmzuxt",
			want: 5,
		},
		{
			name: "ohvhjdml",
			s:    "ohvhjdml",
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
