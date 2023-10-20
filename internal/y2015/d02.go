package y2015

import (
	"sort"
	"strconv"
	"strings"
)

func parseDimensions(present string) []int {
	dims := make([]int, 3)

	for i, dim := range strings.Split(present, "x") {
		num, _ := strconv.Atoi(dim)
		dims[i] = num
	}

	// l, w, h
	return dims
}

func calcPaper(dims []int) int {
	// 2*l*w + 2*w*h + 2*h*l + <area of smallest side>
	sort.Ints(dims)

	return 2*(dims[0]*dims[1]+dims[1]*dims[2]+dims[2]*dims[0]) + (dims[0] * dims[1])
}

func calcRibbon(dims []int) int {
	// <perimeter of smallest side> + l*w*h
	sort.Ints(dims)

	return 2*(dims[0]+dims[1]) + dims[0]*dims[1]*dims[2]
}

func Day02Part1(presents []string) int {
	sum := 0

	for _, present := range presents {
		sum += calcPaper(parseDimensions(present))
	}

	return sum
}

func Day02Part2(presents []string) int {
	sum := 0

	for _, present := range presents {
		sum += calcRibbon(parseDimensions(present))
	}

	return sum
}
