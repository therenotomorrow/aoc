package y2015

import "testing"

func TestDay09Part1(t *testing.T) {
	type args struct {
		distances []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{distances: []string{"London to Dublin = 464", "London to Belfast = 518", "Dublin to Belfast = 141"}}, want: 605},
		{name: "answer", args: args{distances: []string{"Tristram to AlphaCentauri = 34", "Tristram to Snowdin = 100", "Tristram to Tambi = 63", "Tristram to Faerun = 108", "Tristram to Norrath = 111", "Tristram to Straylight = 89", "Tristram to Arbre = 132", "AlphaCentauri to Snowdin = 4", "AlphaCentauri to Tambi = 79", "AlphaCentauri to Faerun = 44", "AlphaCentauri to Norrath = 147", "AlphaCentauri to Straylight = 133", "AlphaCentauri to Arbre = 74", "Snowdin to Tambi = 105", "Snowdin to Faerun = 95", "Snowdin to Norrath = 48", "Snowdin to Straylight = 88", "Snowdin to Arbre = 7", "Tambi to Faerun = 68", "Tambi to Norrath = 134", "Tambi to Straylight = 107", "Tambi to Arbre = 40", "Faerun to Norrath = 11", "Faerun to Straylight = 66", "Faerun to Arbre = 144", "Norrath to Straylight = 115", "Norrath to Arbre = 135", "Straylight to Arbre = 127"}}, want: 251},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day09Part1(tt.args.distances); got != tt.want {
				t.Errorf("Day09Part1() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestDay09Part2(t *testing.T) {
	type args struct {
		distances []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{distances: []string{"London to Dublin = 464", "London to Belfast = 518", "Dublin to Belfast = 141"}}, want: 982},
		{name: "answer", args: args{distances: []string{"Tristram to AlphaCentauri = 34", "Tristram to Snowdin = 100", "Tristram to Tambi = 63", "Tristram to Faerun = 108", "Tristram to Norrath = 111", "Tristram to Straylight = 89", "Tristram to Arbre = 132", "AlphaCentauri to Snowdin = 4", "AlphaCentauri to Tambi = 79", "AlphaCentauri to Faerun = 44", "AlphaCentauri to Norrath = 147", "AlphaCentauri to Straylight = 133", "AlphaCentauri to Arbre = 74", "Snowdin to Tambi = 105", "Snowdin to Faerun = 95", "Snowdin to Norrath = 48", "Snowdin to Straylight = 88", "Snowdin to Arbre = 7", "Tambi to Faerun = 68", "Tambi to Norrath = 134", "Tambi to Straylight = 107", "Tambi to Arbre = 40", "Faerun to Norrath = 11", "Faerun to Straylight = 66", "Faerun to Arbre = 144", "Norrath to Straylight = 115", "Norrath to Arbre = 135", "Straylight to Arbre = 127"}}, want: 898},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day09Part2(tt.args.distances); got != tt.want {
				t.Errorf("Day09Part2() = %v, want = %v", got, tt.want)
			}
		})
	}
}
