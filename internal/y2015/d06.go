package y2015

import (
	"strconv"
	"strings"
)

func newGrid() [][]int {
	grid := make([][]int, 1000)

	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	return grid
}

func parseInstruction(inst string) (cmd string, from []int, to []int) {
	parts := strings.Fields(inst)

	switch parts[0] {
	case "toggle":
		cmd = parts[0]
	case "turn":
		switch parts[1] {
		case "on":
			cmd = "turn on"
		case "off":
			cmd = "turn off"
		}
	}

	parseCoord := func(s string) []int {
		coords := strings.Split(s, ",")

		l, _ := strconv.Atoi(coords[0])
		r, _ := strconv.Atoi(coords[1])

		return []int{l, r}
	}

	from = parseCoord(parts[len(parts)-3])
	to = parseCoord(parts[len(parts)-1])

	return cmd, from, to
}

func cntLights(grid [][]int) int {
	cnt := 0

	for _, row := range grid {
		for _, ceil := range row {
			cnt += ceil
		}
	}

	return cnt
}

// ---- part 1

func turnOn(grid [][]int, from []int, to []int) {
	for i := from[0]; i <= to[0]; i++ {
		for j := from[1]; j <= to[1]; j++ {
			grid[i][j] = 1
		}
	}
}

func turnOff(grid [][]int, from []int, to []int) {
	for i := from[0]; i <= to[0]; i++ {
		for j := from[1]; j <= to[1]; j++ {
			grid[i][j] = 0
		}
	}
}

func toggleLights(grid [][]int, from []int, to []int) {
	for i := from[0]; i <= to[0]; i++ {
		for j := from[1]; j <= to[1]; j++ {
			grid[i][j] ^= 1
		}
	}
}

func Day06Part1(instructions []string) int {
	grid := newGrid()

	for _, inst := range instructions {
		cmd, from, to := parseInstruction(inst)

		switch cmd {
		case "toggle":
			toggleLights(grid, from, to)
		case "turn on":
			turnOn(grid, from, to)
		case "turn off":
			turnOff(grid, from, to)
		}
	}

	return cntLights(grid)
}

// ---- part 2

func brightnessOn(grid [][]int, from []int, to []int) {
	for i := from[0]; i <= to[0]; i++ {
		for j := from[1]; j <= to[1]; j++ {
			grid[i][j]++
		}
	}
}

func brightnessOff(grid [][]int, from []int, to []int) {
	for i := from[0]; i <= to[0]; i++ {
		for j := from[1]; j <= to[1]; j++ {
			if grid[i][j] > 0 {
				grid[i][j]--
			}
		}
	}
}

func toggleBrightness(grid [][]int, from []int, to []int) {
	for i := from[0]; i <= to[0]; i++ {
		for j := from[1]; j <= to[1]; j++ {
			grid[i][j] += 2
		}
	}
}

func Day06Part2(instructions []string) int {
	grid := newGrid()

	for _, inst := range instructions {
		cmd, from, to := parseInstruction(inst)

		switch cmd {
		case "toggle":
			toggleBrightness(grid, from, to)
		case "turn on":
			brightnessOn(grid, from, to)
		case "turn off":
			brightnessOff(grid, from, to)
		}
	}

	return cntLights(grid)
}
