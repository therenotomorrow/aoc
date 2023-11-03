package y2015

import "testing"

func TestDay17Part1(t *testing.T) {
	type args struct {
		target     int
		containers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{target: 25, containers: []int{20, 15, 10, 5, 5}}, want: 4},
		{name: "answer", args: args{target: 150, containers: []int{11, 30, 47, 31, 32, 36, 3, 1, 5, 3, 32, 36, 15, 11, 46, 26, 28, 1, 19, 3}}, want: 4372},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day17Part1(tt.args.target, tt.args.containers); got != tt.want {
				t.Errorf("Day17Part1() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestDay17Part2(t *testing.T) {
	type args struct {
		target     int
		containers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{target: 25, containers: []int{20, 15, 10, 5, 5}}, want: 3},
		{name: "answer", args: args{target: 150, containers: []int{11, 30, 47, 31, 32, 36, 3, 1, 5, 3, 32, 36, 15, 11, 46, 26, 28, 1, 19, 3}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day17Part2(tt.args.target, tt.args.containers); got != tt.want {
				t.Errorf("Day17Part2() = %v, want = %v", got, tt.want)
			}
		})
	}
}
