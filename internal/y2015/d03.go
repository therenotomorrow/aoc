package y2015

import (
	"github.com/therenotomorrow/aoc/pkg/datastruct"
)

type Coord struct {
	x int
	y int
}

func Day03Part1(directions string) int {
	curr := Coord{0, 0}
	coords := datastruct.NewSet()

	coords.Add(curr)

	for _, dir := range directions {
		switch dir {
		case '>':
			curr.x++
		case '^':
			curr.y++
		case '<':
			curr.x--
		case 'v':
			curr.y--
		}

		coords.Add(curr)
	}

	return coords.Size()
}

func Day03Part2(directions string) int {
	santa := &Coord{0, 0}
	robot := &Coord{0, 0}

	curr := santa
	coords := datastruct.NewSet()

	coords.Add(*curr)

	for i, dir := range directions {
		if i%2 == 0 {
			curr = santa
		} else {
			curr = robot
		}

		switch dir {
		case '>':
			curr.x++
		case '^':
			curr.y++
		case '<':
			curr.x--
		case 'v':
			curr.y--
		}

		coords.Add(*curr)
	}

	return coords.Size()
}
