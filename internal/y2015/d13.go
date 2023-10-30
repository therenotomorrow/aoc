package y2015

import (
	"math"
	"strconv"
	"strings"
)

func parseHappiness(happiness []string) (dict map[string]map[string]int, players []string) {
	dict = make(map[string]map[string]int)
	set := make(map[string]struct{})

	for _, happy := range happiness {
		parts := strings.Fields(happy)

		p1 := parts[0]
		p2 := strings.TrimSuffix(parts[10], ".")
		val, _ := strconv.Atoi(parts[3])

		if parts[2] == "lose" {
			val = -val
		}

		if _, ok := dict[p1]; !ok {
			dict[p1] = make(map[string]int)
		}
		if _, ok := dict[p2]; !ok {
			dict[p2] = make(map[string]int)
		}

		dict[p1][p2] = val

		set[p1] = struct{}{}
		set[p2] = struct{}{}
	}

	players = make([]string, 0, len(set))
	for player := range set {
		players = append(players, player)
	}

	return dict, players
}

func permutations(a []string) [][]string {
	choices := make([][]string, 0)

	_permutations(a, func(arr []string) {
		copyA := make([]string, len(a))
		copy(copyA, arr)
		choices = append(choices, copyA)
	}, 0)

	return choices
}

func _permutations(a []string, f func([]string), i int) {
	if i > len(a) {
		f(a)
		return
	}

	_permutations(a, f, i+1)

	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]

		_permutations(a, f, i+1)

		a[i], a[j] = a[j], a[i]
	}
}

func Day13Part1(happiness []string) int {
	dict, players := parseHappiness(happiness)
	maxScore := math.MinInt

	for _, table := range permutations(players) {
		sum := 0

		for i := 0; i < len(table); i++ {
			p1 := table[i]
			p2 := table[(i+1)%len(table)]

			sum += dict[p1][p2]
			sum += dict[p2][p1]
		}

		if sum > maxScore {
			maxScore = sum
		}
	}

	return maxScore
}

func Day13Part2(happiness []string) int {
	dict, players := parseHappiness(happiness)
	maxScore := math.MinInt

	// add myself at the table
	players = append(players, "Kirill")

	dict["Kirill"] = make(map[string]int)

	for _, player := range players {
		dict["Kirill"][player] = 0
	}

	for _, opponents := range dict {
		opponents["Kirill"] = 0
	}

	for _, table := range permutations(players) {
		sum := 0

		for i := 0; i < len(table); i++ {
			p1 := table[i]
			p2 := table[(i+1)%len(table)]

			sum += dict[p1][p2] + dict[p2][p1]
		}

		if sum > maxScore {
			maxScore = sum
		}
	}

	return maxScore
}
