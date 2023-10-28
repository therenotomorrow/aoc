package y2015

import "testing"

func TestDay10Part1(t *testing.T) {
	type args struct {
		seq   string
		times int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{seq: "1", times: 1}, want: 2},
		{name: "smoke 2", args: args{seq: "11", times: 1}, want: 2},
		{name: "smoke 3", args: args{seq: "21", times: 1}, want: 4},
		{name: "smoke 4", args: args{seq: "1211", times: 1}, want: 6},
		{name: "smoke 5", args: args{seq: "111221", times: 1}, want: 6},
		{name: "answer", args: args{seq: "1113222113", times: 40}, want: 252594},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day10Part1(tt.args.seq, tt.args.times); got != tt.want {
				t.Errorf("Day10Part1() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestDay10Part2(t *testing.T) {
	type args struct {
		seq string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "answer", args: args{seq: "1113222113"}, want: 3579328},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day10Part2(tt.args.seq); got != tt.want {
				t.Errorf("Day10Part2() = %v, want = %v", got, tt.want)
			}
		})
	}
}
