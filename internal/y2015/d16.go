package y2015

import (
	"strconv"
	"strings"
)

func parseAunts(aunts []string) []map[string]int {
	sues := make([]map[string]int, len(aunts)+1)

	for _, aunt := range aunts {
		parts := strings.Fields(aunt)

		id, _ := strconv.Atoi(strings.TrimSuffix(parts[1], ":"))

		sues[id] = make(map[string]int)

		for i := 2; i < len(parts); i += 2 {
			key := strings.TrimSuffix(parts[i], ":")
			val, _ := strconv.Atoi(strings.TrimSuffix(parts[i+1], ","))
			sues[id][key] = val
		}
	}

	return sues
}

func Day16Part1(aunts []string) int {
	tmpl := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	trueSue := 0
	maxScore := 0

	for id, sue := range parseAunts(aunts) {
		score := 0

		for k, v := range sue {
			if tmpl[k] == v {
				score++
			}
		}

		if score > maxScore {
			maxScore = score
			trueSue = id
		}
	}

	return trueSue
}

func Day16Part2(aunts []string) int {
	tmpl := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	trueSue := 0
	maxScore := 0

	for id, sue := range parseAunts(aunts) {
		score := 0

		for k, v := range sue {
			switch k {
			case "cats", "trees":
				if v > tmpl[k] {
					score++
				}
			case "pomeranians", "goldfish":
				if v < tmpl[k] {
					score++
				}
			default:
				if tmpl[k] == v {
					score++
				}
			}
		}

		if score > maxScore {
			maxScore = score
			trueSue = id
		}
	}

	return trueSue
}
