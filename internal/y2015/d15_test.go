package y2015

import "testing"

func TestDay15Part1(t *testing.T) {
	type args struct {
		ingredients []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{ingredients: []string{"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8", "Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3"}}, want: 62842880},
		{name: "answer", args: args{ingredients: []string{"Frosting: capacity 4, durability -2, flavor 0, texture 0, calories 5", "Candy: capacity 0, durability 5, flavor -1, texture 0, calories 8", "Butterscotch: capacity -1, durability 0, flavor 5, texture 0, calories 6", "Sugar: capacity 0, durability 0, flavor -2, texture 2, calories 1"}}, want: 18965440},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day15Part1(tt.args.ingredients); got != tt.want {
				t.Errorf("Day15Part1() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestDay15Part2(t *testing.T) {
	type args struct {
		ingredients []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "smoke 1", args: args{ingredients: []string{"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8", "Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3"}}, want: 57600000},
		{name: "answer", args: args{ingredients: []string{"Frosting: capacity 4, durability -2, flavor 0, texture 0, calories 5", "Candy: capacity 0, durability 5, flavor -1, texture 0, calories 8", "Butterscotch: capacity -1, durability 0, flavor 5, texture 0, calories 6", "Sugar: capacity 0, durability 0, flavor -2, texture 2, calories 1"}}, want: 15862900},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day15Part2(tt.args.ingredients); got != tt.want {
				t.Errorf("Day15Part2() = %v, want = %v", got, tt.want)
			}
		})
	}
}
