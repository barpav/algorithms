package basic_calculator_2

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1+2+3",
			args: args{s: "1+2+3"},
			want: 6,
		},
		{
			name: "1+2*3+4*5",
			args: args{s: "1+2*3+4*5"},
			want: 27,
		},
		{
			name: "1+2*3+4*5-6*7-8*9+10*2+11",
			args: args{s: "1+2*3+4*5-6*7-8*9+10*2+11"},
			want: -56,
		},
		{
			name: "3+2*2",
			args: args{s: "3+2*2"},
			want: 7,
		},
		{
			name: " 3/2 ",
			args: args{s: " 3/2 "},
			want: 1,
		},
		{
			name: " 3+5 / 2 ",
			args: args{s: " 3+5 / 2 "},
			want: 5,
		},
		{
			name: " 3+5+6*2*2 - 11-12/3/2  + 7+1  * 3 ",
			args: args{s: " 3+5+6*2*2 - 11-12/3/2  + 7+1  * 3 "},
			want: 29,
		},
		{
			name: "2*3*4",
			args: args{s: "2*3*4"},
			want: 24,
		},
		{
			name: "1+1*2*3*4*5",
			args: args{s: "1+1*2*3*4*5"},
			want: 121,
		},
		{
			name: "1+1*2*3*4*5-1",
			args: args{s: "1+1*2*3*4*5-1"},
			want: 120,
		},
		{
			name: "3-1-1*2*3*4*9/3/3-1",
			args: args{s: "3-1-1*2*3*4*9/3/3-1"},
			want: -23,
		},
		{
			name: "3234-166767 * 3032 * 88 -   0 * 523232 - 2000000 / 2 + 142",
			args: args{s: "3234-166767 * 3032 * 88 -   0 * 523232 - 2000000 / 2 + 142"},
			want: -44497100496,
		},
		{
			name: "0",
			args: args{s: "0"},
			want: 0,
		},
		{
			name: "2*3+4",
			args: args{s: "2*3+4"},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
