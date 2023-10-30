// Package day09: https://adventofcode.com/2022/day/9
package day09

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction = string

type Step = int

const (
	Up    = "U"
	Down  = "D"
	Left  = "L"
	Right = "R"
)

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

var visited = make(map[Point]struct{})

type Point struct {
	X int
	Y int
}

type Rope struct {
	Head *Point
	Tail *Point

	Knots []*Point
}

func NewRope(l int) Rope {
	points := make([]*Point, l)

	for i := 0; i < l; i++ {
		points[i] = &Point{}
	}

	return Rope{Knots: points}
}

func (r *Rope) Move(d Direction, steps Step) {
	for step := 0; step < steps; step++ {
		r.Head = r.Knots[0]

		switch d {
		case Up:
			r.Head.Y++
		case Down:
			r.Head.Y--
		case Left:
			r.Head.X--
		case Right:
			r.Head.X++
		}

		for i := 0; i < len(r.Knots)-1; i++ {
			r.Head = r.Knots[i]
			r.Tail = r.Knots[i+1]

			r.MoveTail()
		}

		visited[*r.Tail] = struct{}{}
	}
}

func (r *Rope) MoveTail() {
	switch {
	case r.Head.X == r.Tail.X && r.Head.Y == r.Tail.Y:
		return
	case r.Head.X == r.Tail.X:
		r.moveTailV()
	case r.Head.Y == r.Tail.Y:
		r.moveTailH()
	default:
		r.moveTailD()
	}
}

func (r *Rope) moveTailV() {
	if abs(r.Head.Y-r.Tail.Y) < 2 {
		return
	}

	if r.Head.Y > r.Tail.Y {
		r.Tail.Y++
	} else {
		r.Tail.Y--
	}
}

func (r *Rope) moveTailH() {
	if abs(r.Head.X-r.Tail.X) < 2 {
		return
	}

	if r.Head.X > r.Tail.X {
		r.Tail.X++
	} else {
		r.Tail.X--
	}
}

func (r *Rope) moveTailD() {
	if abs(r.Head.X-r.Tail.X) < 2 && abs(r.Head.Y-r.Tail.Y) < 2 {
		return
	}

	if abs(r.Head.X-r.Tail.X) < 2 {
		r.moveTailV()

		if r.Head.X > r.Tail.X {
			r.Tail.X++
		} else {
			r.Tail.X--
		}
		return
	}

	if abs(r.Head.Y-r.Tail.Y) < 2 {
		r.moveTailH()

		if r.Head.Y > r.Tail.Y {
			r.Tail.Y++
		} else {
			r.Tail.Y--
		}
		return
	}

	r.moveTailH()
	r.moveTailV()
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scan := bufio.NewScanner(f)

	var moves []string

	for scan.Scan() {
		moves = append(moves, scan.Text())
	}

	bridge := NewRope(10)

	for _, move := range moves {
		tmp := strings.Fields(move)

		d := tmp[0]
		step, _ := strconv.Atoi(tmp[1])

		bridge.Move(d, step)
	}

	fmt.Println(len(visited))
}
