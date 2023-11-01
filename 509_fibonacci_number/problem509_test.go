package problem509

import "testing"

func Test_fib(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "n = 0",
			n:    0,
			want: 0,
		},
		{
			name: "n = 1",
			n:    1,
			want: 1,
		},
		{
			name: "n = 2",
			n:    2,
			want: 1,
		},
		{
			name: "n = 3",
			n:    3,
			want: 2,
		},
		{
			name: "n = 4",
			n:    4,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
