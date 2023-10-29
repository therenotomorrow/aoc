package y2015

import "testing"

func TestDay11Part1(t *testing.T) {
	type args struct {
		prev string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "smoke 1", args: args{prev: "abcdefgh"}, want: "abcdffaa"},
		{name: "smoke 2", args: args{prev: "ghijklmn"}, want: "ghjaabcc"},
		{name: "answer", args: args{prev: "hepxcrrq"}, want: "hepxxyzz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day11Part1(tt.args.prev); got != tt.want {
				t.Errorf("Day11Part1() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestDay11Part2(t *testing.T) {
	type args struct {
		prev string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "answer", args: args{prev: "hepxxyzz"}, want: "heqaabcc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day11Part2(tt.args.prev); got != tt.want {
				t.Errorf("Day11Part2() = %v, want = %v", got, tt.want)
			}
		})
	}
}
