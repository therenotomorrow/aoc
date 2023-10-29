package y2015

import (
	"math"
	"strconv"
	"strings"
)

func travel(dict map[string]map[string]int, start string, predicate func(int, int) bool) int {
	total := 0
	visited := make(map[string]struct{})

	for start != "" {
		next := ""
		last := 0

		for choice, interval := range dict[start] {
			if _, ok := visited[choice]; ok {
				continue
			}

			if predicate(interval, last) {
				last = interval
				next = choice
			}
		}

		// we travel all the cities
		if next == "" {
			break
		}

		total += last
		visited[start] = struct{}{}
		start = next
	}

	return total
}

func parseDistances(distances []string) (dict map[string]map[string]int, cities map[string]struct{}) {
	dict = make(map[string]map[string]int)
	cities = make(map[string]struct{})

	for _, distance := range distances {
		parts := strings.Fields(distance)
		from := parts[0]
		to := parts[2]
		interval, _ := strconv.Atoi(parts[4])

		if _, ok := dict[from]; !ok {
			dict[from] = make(map[string]int)
		}

		if _, ok := dict[to]; !ok {
			dict[to] = make(map[string]int)
		}

		dict[from][to] = interval
		dict[to][from] = interval

		cities[from] = struct{}{}
		cities[to] = struct{}{}
	}

	return dict, cities
}

func Day09Part1(distances []string) int {
	minDist := math.MaxInt
	dict, cities := parseDistances(distances)

	for city := range cities {
		minDist = min(minDist, travel(dict, city, func(a int, b int) bool {
			if b == 0 {
				b = math.MaxInt
			}
			return a < b
		}))
	}

	return minDist
}

func Day09Part2(distances []string) int {
	maxDist := math.MinInt
	dict, cities := parseDistances(distances)

	for city := range cities {
		maxDist = max(maxDist, travel(dict, city, func(a int, b int) bool {
			if b == 0 {
				b = math.MinInt
			}
			return a > b
		}))
	}

	return maxDist
}
