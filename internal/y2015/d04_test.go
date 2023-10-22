package y2015

import "testing"

func TestDay04Part1(t *testing.T) {
	type args struct {
		salt string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{salt: "abcdef"}, want: 609043},
		{name: "smoke 2", args: args{salt: "pqrstuv"}, want: 1048970},
		{name: "answer", args: args{salt: "iwrupvqb"}, want: 346386},
		{name: "impossible", args: args{salt: ""}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day04Part1(tt.args.salt); got != tt.want {
				t.Errorf("Day04Part1() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestDay04Part2(t *testing.T) {
	type args struct {
		salt string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "answer", args: args{salt: "iwrupvqb"}, want: 9958218},
		{name: "impossible", args: args{salt: ""}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day04Part2(tt.args.salt); got != tt.want {
				t.Errorf("Day04Part2() = %v, want = %v", got, tt.want)
			}
		})
	}
}
