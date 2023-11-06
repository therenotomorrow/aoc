package y2015

import "testing"

func TestDay20Part1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 9}, want: 1},
		{name: "answer", args: args{n: 34000000}, want: 786240},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day20Part1(tt.args.n); got != tt.want {
				t.Errorf("Day20Part1() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestDay20Part2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{n: 9}, want: 1},
		{name: "answer", args: args{n: 34000000}, want: 831600},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day20Part2(tt.args.n); got != tt.want {
				t.Errorf("Day20Part2() = %v, want = %v", got, tt.want)
			}
		})
	}
}
