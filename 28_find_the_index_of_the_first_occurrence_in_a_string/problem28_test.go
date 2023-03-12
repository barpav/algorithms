package problem28

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "'sad' in: 'sadbutsad'",
			args: args{haystack: "sadbutsad", needle: "sad"},
			want: 0,
		},
		{
			name: "'leeto' in: 'leetcode'",
			args: args{haystack: "leetcode", needle: "leeto"},
			want: -1,
		},
		{
			name: "'issip' in: 'mississippi'",
			args: args{haystack: "mississippi", needle: "issip"},
			want: 4,
		},
		{
			name: "'a' in: 'a'",
			args: args{haystack: "a", needle: "a"},
			want: 0,
		},
		{
			name: "'c' in: 'abc'",
			args: args{haystack: "abc", needle: "c"},
			want: 2,
		},
		{
			name: "'7 feet' in: 'The temple measured approximately...'",
			args: args{haystack: "The temple measured approximately 90 ft x 22 ft, with the supporting platform extending 7 ft beyond the row of columns. The columns are approximately 12 ft high and 1 ft in diameter, standing 7 feet apart.", needle: "7 feet"},
			want: 192,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
