package y2015

import "testing"

func TestDay14Part1(t *testing.T) {
	type args struct {
		deer []string
		time int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{deer: []string{"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.", "Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds."}, time: 1000}, want: 1120},
		{name: "answer", args: args{deer: []string{"Dancer can fly 27 km/s for 5 seconds, but then must rest for 132 seconds.", "Cupid can fly 22 km/s for 2 seconds, but then must rest for 41 seconds.", "Rudolph can fly 11 km/s for 5 seconds, but then must rest for 48 seconds.", "Donner can fly 28 km/s for 5 seconds, but then must rest for 134 seconds.", "Dasher can fly 4 km/s for 16 seconds, but then must rest for 55 seconds.", "Blitzen can fly 14 km/s for 3 seconds, but then must rest for 38 seconds.", "Prancer can fly 3 km/s for 21 seconds, but then must rest for 40 seconds.", "Comet can fly 18 km/s for 6 seconds, but then must rest for 103 seconds.", "Vixen can fly 18 km/s for 5 seconds, but then must rest for 84 seconds."}, time: 2503}, want: 2640},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day14Part1(tt.args.deer, tt.args.time); got != tt.want {
				t.Errorf("Day14Part1() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestDay14Part2(t *testing.T) {
	type args struct {
		deer []string
		time int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{deer: []string{"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.", "Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds."}, time: 1000}, want: 689},
		{name: "answer", args: args{deer: []string{"Dancer can fly 27 km/s for 5 seconds, but then must rest for 132 seconds.", "Cupid can fly 22 km/s for 2 seconds, but then must rest for 41 seconds.", "Rudolph can fly 11 km/s for 5 seconds, but then must rest for 48 seconds.", "Donner can fly 28 km/s for 5 seconds, but then must rest for 134 seconds.", "Dasher can fly 4 km/s for 16 seconds, but then must rest for 55 seconds.", "Blitzen can fly 14 km/s for 3 seconds, but then must rest for 38 seconds.", "Prancer can fly 3 km/s for 21 seconds, but then must rest for 40 seconds.", "Comet can fly 18 km/s for 6 seconds, but then must rest for 103 seconds.", "Vixen can fly 18 km/s for 5 seconds, but then must rest for 84 seconds."}, time: 2503}, want: 1102},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day14Part2(tt.args.deer, tt.args.time); got != tt.want {
				t.Errorf("Day14Part2() = %v, want = %v", got, tt.want)
			}
		})
	}
}
